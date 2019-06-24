// Countdown
const timeUnits = [
  ['Y', 1000 * 60 * 60 * 24 * 365], // years
  ['M', 1000 * 60 * 60 * 24 * 30], // months
  ['D', 1000 * 60 * 60 * 24], // days
  ['H', 1000 * 60 * 60], // hours
  ['m', 1000 * 60], // minutes
  ['s', 1000], // seconds
  ['S', 1], // million seconds
];

export const formatTimeStr = (duration, format = 'mm:ss') => {
  let leftDuration = Math.abs(duration);

  const formatAbs =  timeUnits.reduce((current, [name, unit]) => {
    if (current.indexOf(name) !== -1) {
      const value = Math.floor(leftDuration / unit);
      leftDuration -= value * unit;
      return current.replace(new RegExp(`${name}+`, 'g'), (match) => {
        const result = value.toString();
        return '0'.repeat(Math.max(match.length - result.length, 0)) + result;
      });
    }
    return current;
  }, format);
  return duration < 0 ? '-'+formatAbs : formatAbs
};

// 22:30:59 => 
export const str2MillionSeconds = (str) => {
  const [seconds, minutes, hours = 0 ] = str.split(':').reverse();
  return ((((hours * 60) + minutes) * 60) + Number(seconds)) * 1000
}