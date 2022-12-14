import "cross-fetch/polyfill";
import { Fetcher } from "openapi-typescript-fetch";
import { env } from "../env";

import type { paths } from "../generated/schema";

// if (env.NORMAL_TOKEN == "" || env.STUDENT_TOKEN == "") {
//   throw new Error("NORMAL_TOKEN or STUDENT_TOKEN is empty");
// }

const InternalFetcher = (XInternalAuthorization: string) => {
  const internalFetcher = Fetcher.for<paths>();

  internalFetcher.configure({
    baseUrl: env.BASEURL,
    init: {
      headers: {
        "X-InternalAuthorization": XInternalAuthorization,
      },
    },
  });

  return internalFetcher;
};

export const SudoFetcher = InternalFetcher(
  env.SUDO_TOKEN,
);

export const ReadFetcher = InternalFetcher(
    env.READ_TOKEN,
)

export const WriteFetcher = InternalFetcher(
    env.WRITE_TOKEN,
)