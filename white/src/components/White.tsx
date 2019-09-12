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

import { connect } from '../utils/rtc'

const canvasWidth = 1080 * 3;
const canvasHeight = 1080 * 3;

interface WhiteProps {}

const useCanvasCtx = (): [
  (canvas: HTMLCanvasElement) => void,
  React.MutableRefObject<CanvasRenderingContext2D | null>,
  ClientRect
] => {
  const [rect, setRect] = useState<ClientRect>();
  const ctxRef = useRef<CanvasRenderingContext2D | null>(null);
  const ref = useCallback((canvas: HTMLCanvasElement) => {
    setRect(canvas.getBoundingClientRect());
    ctxRef.current = canvas.getContext("2d");
  }, []);
  return [ref, ctxRef, rect as ClientRect];
};

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
  style: CtxStyle;
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
  stopHandler?: (p: DrawPoint) => void
) => {
  const start = (event: MouseEvent | TouchEvent) => {
    drawing = true;
    startHandler(event2Point(event));
  };
  const move = (event: MouseEvent | TouchEvent) => {
    if (drawing) {
      event.stopPropagation();
      moveHandler(event2Point(event));
    }
  };
  const stop = (event: MouseEvent | TouchEvent) => {
    drawing = false;
    stopHandler && stopHandler(event2Point(event));
  };
  return { start, move, stop };
};

const drawLineHandler = (ctx: CanvasRenderingContext2D, draw: Draw) => {
  ctx.strokeStyle = draw.style;
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

  const [ref, ctxRef, rect] = useCanvasCtx();
  const [style, setStyle] = useState<CtxStyle>("");
  const [dialOpen, setDialOpen] = useState<boolean>(false);
  const drawsRef = useRef<Draws>([]);
  const colorRef = useRef<HTMLInputElement>(null);


  const drawHandler = useCallback((line: Draw) => {
    const ctx = ctxRef.current as CanvasRenderingContext2D;
    drawLineHandler(ctx, line);
  }, []);

  // useEffect(() => {
  //   console.log('useEffect')
  //   const ctx = ctxRef.current as CanvasRenderingContext2D;
  //   ctx.clearRect(0,0,canvasWidth, canvasHeight)
  //   drawsRef.current.forEach(drawHandler)
  // });

  const setCurrentPoint = (p: DrawPoint) => {
    const current = drawsRef.current[drawsRef.current.length - 1];
    const point: DrawPoint = {
      x: Math.round(((p.x - rect.left) * canvasWidth) / rect.width),
      y: Math.round(((p.y - rect.top) * canvasHeight) / rect.height)
    };
    current.points.push(point);
    drawHandler(current);
  };


  const moveHandlerThrottle = throttle(point => setCurrentPoint(point), 16);

  const drawLine = draw(
    point => {
      drawsRef.current.push({ style, points: [] });
      moveHandlerThrottle(point);
    },
    point => {
      moveHandlerThrottle(point);
    }
  );

  const actions = [
    { icon: <SaveIcon />, name: "Save" },
    // { icon: <PrintIcon />, name: "Print" },
    // { icon: <ShareIcon />, name: "Share" },
    // { icon: <DeleteIcon />, name: "Delete" },
    {
      icon: <InvertColors />, 
      name: "color",
      hanler() {
        (colorRef.current as HTMLInputElement).click()
        setDialOpen(false)
      }
    }
  ];

  // web rtc
  useEffect(()=>{
    connect()
  },[])

  return (
    <div>
      <canvas
        className={classes.root}
        ref={ref}
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
        ref={colorRef}
        type="color"
        hidden
        onChange={e => setStyle(e.target.value)}
      />
    </div>
  );
};

export default White;
