import React, {
  useRef,
  useCallback,
  useState,
  useEffect,
  MouseEvent,
  TouchEvent,
  Touch
} from "react";
import { makeStyles } from "@material-ui/core/styles";
import throttle from "lodash/throttle";
import SpeedDial from "@material-ui/lab/SpeedDial";
import SpeedDialIcon from "@material-ui/lab/SpeedDialIcon";
import SpeedDialAction from "@material-ui/lab/SpeedDialAction";
import InvertColors from "@material-ui/icons/InvertColors";
import SaveIcon from "@material-ui/icons/Save";

import { connect, Connection } from '../utils/rtc'
import { Message } from '../utils/message'
import { decode } from '../utils/textEncoder'

// TODO: remove this
declare global {
  interface Window {
    conn: any
  }
}

const canvasWidth = 1080 * 3;
const canvasHeight = 1080 * 3;

interface WhiteProps {}


const useStyles = makeStyles(theme => ({
  root: {
    width: "100%",
    height: "100vh",
    display: "block",
    touchAction: "none"
  },
  speedDial: {
    position: "absolute",
    bottom: theme.spacing(2),
    right: theme.spacing(3)
  }
}));

interface DrawPoint {
  x: number;
  y: number;
}

type CtxStyle = string | CanvasGradient | CanvasPattern;
interface Draw {
  color: CtxStyle;
  points: DrawPoint[];
}
type Draws = Draw[];

const event2Point = (event: MouseEvent | TouchEvent): DrawPoint => {
  const point = {
    x: 0,
    y: 0
  };
  let value: Touch | MouseEvent;
  switch (event.type) {
    case "touchstart":
    case "touchmove":
    case "touchend":
      value = (event as TouchEvent).changedTouches[0];
      break;
    case "mousedown":
    case "mousemove":
    case "mouseup":
      value = event as MouseEvent;
      break;
    default:
      value = event as MouseEvent;
      break;
  }

  point.x = value.clientX;
  point.y = value.clientY;
  return point;
};
let drawing = false;
const draw = (
  startHandler: (p: DrawPoint) => void,
  moveHandler: (p: DrawPoint) => void,
  stopHandler: (p: DrawPoint) => void = ()=>{}
) => {
  const moveHandlerThrottle = throttle(moveHandler, 16);

  const start = (event: MouseEvent | TouchEvent) => {
    drawing = true;
    startHandler(event2Point(event));
  };
  const move = (event: MouseEvent | TouchEvent) => {
    if (drawing) {
      event.stopPropagation();
      moveHandlerThrottle(event2Point(event));
    }
  };
  const stop = (event: MouseEvent | TouchEvent) => {
    drawing = false;
    stopHandler(event2Point(event));
  };
  return { start, move, stop };
};

const drawLineHandler = (ctx: CanvasRenderingContext2D, draw: Draw) => {
  ctx.strokeStyle = draw.color;
  ctx.lineWidth = 15;
  ctx.beginPath();
  ctx.moveTo(draw.points[0].x, draw.points[0].y);
  draw.points.forEach(line => {
    ctx.lineTo(line.x, line.y);
  });
  ctx.stroke();
};

const White: React.FC<WhiteProps> = props => {
  const classes = useStyles();
  console.log("run function components");

  const cvsRef = useRef<HTMLCanvasElement >(null)
  const colorRef = useRef<CtxStyle>("");
  const [dialOpen, setDialOpen] = useState<boolean>(false);
  const drawsRef = useRef<Draws>([]);
  const colorInputRef = useRef<HTMLInputElement>(null);
  const connRef = useRef<Connection>()


  // useEffect(() => {
  //   console.log('useEffect')
  //   const ctx = ctxRef.current as CanvasRenderingContext2D;
  //   ctx.clearRect(0,0,canvasWidth, canvasHeight)
  //   drawsRef.current.forEach(drawHandler)
  // });

  const setCurrentPoint = (p: DrawPoint) => {
    const rect = (cvsRef.current as HTMLCanvasElement).getBoundingClientRect()
    const current = drawsRef.current[drawsRef.current.length - 1];
    const point: DrawPoint = {
      x: Math.round(((p.x - rect.left) * canvasWidth) / rect.width),
      y: Math.round(((p.y - rect.top) * canvasHeight) / rect.height)
    };
    current.points.push(point);
    const ctx = (cvsRef.current as HTMLCanvasElement).getContext("2d") as CanvasRenderingContext2D;
    drawLineHandler(ctx, current);
  };

  const sendPointMsg = (data: any) => {
    const conn = connRef.current
    if(conn) {
      const msg = {
        ...data,
        uid: conn.uid,
       }
      conn.dataChannelsSend(msg)
    }
  }

  const drawLine = draw(
    point => {
      drawsRef.current.push({ color: colorRef.current, points: [] });
      setCurrentPoint(point);
      sendPointMsg({kind: 'pointStart', value: point})
    },
    point => {
      setCurrentPoint(point);
      sendPointMsg({kind: 'pointMove', value: point})
    }
  );

  const colorChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const colorValue = event.target.value
    colorRef.current = colorValue
    sendPointMsg({kind: 'colorChange', value: colorValue})
  }

  const actions = [
    { icon: <SaveIcon />, name: "Save" },
    // { icon: <PrintIcon />, name: "Print" },
    // { icon: <ShareIcon />, name: "Share" },
    // { icon: <DeleteIcon />, name: "Delete" },
    {
      icon: <InvertColors />, 
      name: "color",
      hanler() {
        (colorInputRef.current as HTMLInputElement).click()
        setDialOpen(false)
      }
    }
  ];

  // web rtc
  useEffect(()=>{
    const datachannelmessage = (event: Event) => {
      const dataMsg = ((event as CustomEvent).detail as MessageEvent)
      const decodeMsg = decode(dataMsg.data as Uint8Array)
      console.log(decodeMsg)
      switch(decodeMsg.kind) {
        case 'pointStart':
          drawsRef.current.push({ color: colorRef.current, points: [] });
          setCurrentPoint(decodeMsg.value);
          break;
        case 'pointMove':
          setCurrentPoint(decodeMsg.value);
          break;
        case 'colorChange':
          colorRef.current = decodeMsg.value
          break;
      }
    }
    const init = async () => {
      const conn = await connect();
      (window.conn as any) = connRef.current = conn // TODO: remove this
      conn.addEventListener('datachannelmessage', datachannelmessage)
    }
    init()
    return () => {
      const conn = connRef.current
      if (conn) {
        conn.removeEventListener('datachannelmessage', datachannelmessage)
        conn.handleClose()
      }
    }
  },[])

  return (
    <div>
      <canvas
        className={classes.root}
        ref={cvsRef}
        height={canvasHeight}
        width={canvasWidth}
        onMouseDown={drawLine.start}
        onMouseMove={drawLine.move}
        onMouseUp={drawLine.stop}
        onTouchStart={drawLine.start}
        onTouchMove={drawLine.move}
        onTouchEnd={drawLine.stop}
        // onPointerDown={drawLine.start}
        // onPointerMove={drawLine.move}
        // onPointerUp={drawLine.stop}
      />
      <SpeedDial
        ariaLabel="SpeedDial"
        className={classes.speedDial}
        // hidden={dialOpen}
        icon={<SpeedDialIcon />}
        onClick={() => setDialOpen(o=>!o)}
        onBlur={() => setDialOpen(false)}
        onClose={() => setDialOpen(false)}
        onFocus={() => setDialOpen(true)}
        onMouseEnter={() => setDialOpen(true)}
        onMouseLeave={() => setDialOpen(false)}
        open={dialOpen}
        // direction={direction}
      >
        {actions.map(action => (
          <SpeedDialAction
            key={action.name}
            icon={action.icon}
            tooltipTitle={action.name}
            onClick={action.hanler || console.log}
          />
        ))}
      </SpeedDial>
      <input
        ref={colorInputRef}
        type="color"
        hidden
        onChange={colorChange}
      />
    </div>
  );
};

export default White;
