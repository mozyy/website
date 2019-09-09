
export interface Message {
  kind: 'join' | 'leave' | 'event',
  value: any
}

