import React, { useRef, useCallback, useState, useEffect, MouseEvent, TouchEvent, Touch } from "react";
import { makeStyles } from '@material-ui/styles'


const canvasWidth = 1920 * 3
const canvasHeight = 1080 * 3

interface WhiteProps {
}

const useCanvasCtx = ():[(canvas: HTMLCanvasElement) => void, React.MutableRefObject<CanvasRenderingContext2D | null>, ClientRect] => {
  const [rect, setRect] = useState<ClientRect|null>(null);
  const ctxRef = useRef<CanvasRenderingContext2D | null>(null)
  const ref = useCallback((canvas: HTMLCanvasElement)=>{
    setRect(canvas.getBoundingClientRect());
    ctxRef.current = canvas.getContext('2d')
  },[])
  return [ref, ctxRef, rect as ClientRect]
}

const useStyles = makeStyles({
  root: {
    width: '100%',
    height: '100vh',
    display: 'block',
  }
})

interface DrawPoint {
  x: number,
  y: number,
}
interface Draw {
  style: string | CanvasGradient | CanvasPattern,
  points: DrawPoint[]
}
type Draws = Draw[]

const event2Point = (event: MouseEvent|TouchEvent):DrawPoint => {
  const point = {
    x: 0,
    y: 0,
  }
  let value: Touch|MouseEvent;
  switch(event.type) {
    case 'touchstart':
    case 'touchmove':
      value = (event as TouchEvent).targetTouches[0]
      break;
    case 'touchend':
      value = (event as TouchEvent).changedTouches[0]
      break;
    case 'mousedown':
    case 'mousemove':
    case 'mouseup':
      value = (event as MouseEvent)
      break;
    default:
      return point
  }

  point.x = value.clientX
  point.y = value.clientY
  return point
}
let drawing = false;
const draw = (startHandler: (p: DrawPoint)=>void, moveHandler:(p: DrawPoint)=>void, stopHandler: (p: DrawPoint)=>void)=> {
  const start = (event: MouseEvent|TouchEvent) => {
    drawing = true;
    startHandler(event2Point(event))
  }
  const move = (event: MouseEvent|TouchEvent) => {
    if (drawing) {
      moveHandler(event2Point(event))
    }
  }
  const stop = (event: MouseEvent|TouchEvent) => {
    drawing = false;
    stopHandler(event2Point(event))
  }
  return { start, move, stop }
}


const White: React.FC<WhiteProps> = (props)=> {
  const classes = useStyles()
  const [height,setHeight] = useState(100)
  const [ref, ctxRef, rect] = useCanvasCtx()
  const [draws, setDraws] = useState<Draws>([])
  const [currentDraw, setCurrentDraw] = useState<Draw>({style: 'red', points:[]})

  useEffect(()=>{
    const ctx = ctxRef.current as CanvasRenderingContext2D
    ctx.clearRect(0,0, canvasWidth, canvasHeight)

    const drawLineHandler = (draw:Draw)=> {
      ctx.strokeStyle = draw.style
      ctx.lineWidth = 15
      ctx.beginPath()
      draw.points.forEach((line,index)=> {
        if (index === 0 ) {
          ctx.moveTo(line.x, line.y)
          ctx.lineTo(line.x, line.y)
        } else {
          ctx.lineTo(line.x, line.y)
        }
      })
      ctx.stroke()
    }
    draws.forEach(drawLineHandler)
    drawLineHandler(currentDraw)
  },[draws, currentDraw])

  const drawLine = draw((d)=>{
    const point: DrawPoint = {
      x: d.x * canvasWidth  /rect.width,
      y: d.y * canvasHeight / rect.height
    }
    setCurrentDraw(d=>({
      style:d.style,
      points: [...d.points, point]
    }))
    console.log(rect)
  }, (d)=>{
    const point: DrawPoint = {
      x: d.x * canvasWidth  /rect.width,
      y: d.y * canvasHeight / rect.height
    }
    setCurrentDraw(d=>({
      style:d.style,
      points: [...d.points, point]
    }))
  },()=>{
    setDraws(d=>[...d,currentDraw])
    setCurrentDraw({style: 'red', points:[]})
  })


  return (
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
    >
      {/* <SpeedDial
            ariaLabel="SpeedDial"
            className={speedDialClassName}
            hidden={hidden}
            icon={<SpeedDialIcon />}
            onBlur={this.handleClose}
            onClick={this.handleClick}
            onClose={this.handleClose}
            onFocus={this.handleOpen}
            onMouseEnter={this.handleOpen}
            onMouseLeave={this.handleClose}
            open={open}
            direction={direction}
          >
            {actions.map(action => (
              <SpeedDialAction
                key={action.name}
                icon={action.icon}
                tooltipTitle={action.name}
                onClick={this.handleClick}
              />
            ))}
          </SpeedDial> */}
    </canvas>
  )
}

export default White