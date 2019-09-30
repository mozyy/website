
if (process.env.NODE_ENV === 'development' && navigator.userAgent.match (/ Android/i)) {
  // import('vconsole').then(module=> new module.default())
}

// fix react typesript isolatedMoudles is true
export const noop = () => {}