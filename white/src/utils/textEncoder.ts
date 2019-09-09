import {Message} from './message'


const encoding = 'utf-8';

export const encode = (data: Message) => {
  const dataStr = JSON.stringify(data);
  return new TextEncoder().encode(dataStr);
}

export const decode = (arr: Uint8Array): Message => {
  const dataStr = new TextDecoder().decode(arr)
  return JSON.parse(dataStr)
}