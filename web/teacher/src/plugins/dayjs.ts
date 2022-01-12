import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import timezone from 'dayjs/plugin/timezone'
import 'dayjs/locale/ja'

dayjs.extend(utc)
dayjs.extend(timezone)

dayjs.locale('ja')
dayjs.tz.setDefault('Asia/Tokyo')

export default dayjs
