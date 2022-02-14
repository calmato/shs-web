import dayjs, { extend, locale } from 'dayjs'
import utc from 'dayjs/plugin/utc'
import timezone from 'dayjs/plugin/timezone'
import 'dayjs/locale/ja'

extend(utc)
extend(timezone)

locale('ja')
dayjs.tz.setDefault('Asia/Tokyo')

export default dayjs
