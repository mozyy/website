import config from '../../docker/config.json'

declare var process : {
  env :{
    NODE_ENV: 'development' | 'production'
  }
}

interface ConfigEnv {
  websocket: string,
}

export const configEnv: ConfigEnv = config[process.env.NODE_ENV]


export const URLWs = configEnv.websocket