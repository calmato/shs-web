import http from 'k6/http';
import { check, sleep } from 'k6';

// @see: https://k6.io/docs/using-k6/options/
export const options = {
  // @see: https://k6.io/docs/using-k6/scenarios
  scenarios: {
    teacher_5vus: {
      // @see: https://k6.io/docs/using-k6/scenarios/#executors
      executor: 'shared-iterations',
      startTime: '0s',
      gracefulStop: '10s',
      exec: 'default',
      env: {
        API_BASE_URL: 'https://teacher.shs-web-stg.calmato.jp',
        USER_EMAIL: 'admin@calmato.jp',
        USER_PASSWORD: '12345678',
        TOKEN: 'eyJhbGciOiJSUzI1NiIsImtpZCI6ImYyNGYzMTQ4MTk3ZWNlYTUyOTE3YzNmMTgzOGFiNWQ0ODg3ZWEwNzYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vc2hzLXdlYi1kZXYtYTliNjQiLCJhdWQiOiJzaHMtd2ViLWRldi1hOWI2NCIsImF1dGhfdGltZSI6MTY0NDU5NDc4NywidXNlcl9pZCI6ImNuZ3hLMlliUWtpVWZSVWNwOHpTZXQiLCJzdWIiOiJjbmd4SzJZYlFraVVmUlVjcDh6U2V0IiwiaWF0IjoxNjQ0NjQ5OTM2LCJleHAiOjE2NDQ2NTM1MzYsImVtYWlsIjoiYWRtaW5AY2FsbWF0by5qcCIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJhZG1pbkBjYWxtYXRvLmpwIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifX0.G9aLpcKoJevjiEJSOgjV5a5ttLJydO3GSbyGLELmdT1ZyMWv-MNLB0pus_RWVH0lNiMMDGiZW7cgi8WhWoy4f0ID57jPYguPElSO5p60NB3QnS3P9TYXC8-F00YTIUw8UHu6Aa9ORX24Om12gJJJDfBDuyXtJWsfz4cnOinHrUg_6rloZYIg4mMz5bbg9xkDZRsJEe9Ab2L8VJuCW0P9EB5PgZU7bqaAxmsg8OoLDBRccrO0D62q-r6BeSLbStUrXa_bWmtNC8Ev4s_m7f_YCL5vsx4LbZWRtz17KC9beNx9xGjz14dzg5OIwbRzfmT_A1W77YAnYSCgSKRk4yNvPQ',
      },
      vus: 5,
      iterations: 30,
      maxDuration: '5m00s',
    },
  },
};

export function setup() {}

export default function () {
  const headers = {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${__ENV.TOKEN}`,
    'User-Agent': 'k6',
  };

  /**
   * 認証情報取得API (/v1/me)
   */
  let res = http.get(`${__ENV.API_BASE_URL}/v1/me`, { headers })
  check(res, { 'status was 200': (r) => r.status === 200 });
  sleep(1);

  const teacherId = JSON.parse(res.body || '').id

  /**
   * 授業科目一覧取得API (/v1/subjects)
   */
  res = http.get(`${__ENV.API_BASE_URL}/v1/subjects`, { headers })
  check(res, { 'status was 200': (r) => r.status === 200 });
  sleep(1);

  /**
   * トップ情報取得API (/v1/lessons)
   */
  res = http.get(`${__ENV.API_BASE_URL}/v1/lessons`, { headers })
  check(res, { 'status was 200': (r) => r.status === 200 });
  sleep(1);

  /**
   * 講師一覧取得API (/v1/teachers)
   */
  res = http.get(`${__ENV.API_BASE_URL}/v1/teachers`, { headers })
  check(res, { 'status was 200': (r) => r.status === 200 });
  sleep(1);

  /**
   * シフト募集一覧取得API (/v1/teachers/${teacherId}/submissions)
   */
  res = http.get(`${__ENV.API_BASE_URL}/v1/teachers/${teacherId}/submissions`, { headers })
  check(res, { 'status was 200': (r) => r.status === 200 });
  sleep(1);

  /**
   * シフト詳細取得API (/v1/shifts/${shiftId})
   */
  res = http.get(`${__ENV.API_BASE_URL}/v1/shifts/1`, { headers })
  check(res, { 'status was 200': (r) => r.status === 200 });
  sleep(1);
}
