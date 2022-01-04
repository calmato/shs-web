import dayjs, { extend } from 'dayjs'
import utc from 'dayjs/plugin/utc'
import timezone from 'dayjs/plugin/timezone'

extend(utc)
extend(timezone)

dayjs.tz.setDefault('Asia/Tokyo')

export default dayjs
