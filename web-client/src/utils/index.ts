export function tibianTime(date?: string) {
  if (!date) {
    return 'never'
  }
  
  const d = new Date(date)
  const ssTime = new Date(date)
  ssTime.setHours(10, 0, 0, 0)
  const diff = Math.floor((d.getTime() - ssTime.getTime()) / 1000 / 60 / 60)
  // TODO: handle +/- properly and use german timezone for sever save
  return `${d.toLocaleDateString()} ${d.toLocaleTimeString()} ss+${diff}`
}

export function toCamelCase(str: string): string {
  return str.replace(/([-_][a-z])/g, (group) =>
    group.toUpperCase().replace('-', '').replace('_', '')
  )
}
