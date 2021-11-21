import { HelloResponse } from "~/types/api/v1";

export const hello: { [key: string]: HelloResponse } = {
  '/v1/hello': {
    message: "test message"
  },
}
