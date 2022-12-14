import "cross-fetch/polyfill";
import { Fetcher } from "openapi-typescript-fetch";
import { env } from "../env";

import type { paths } from "../generated/schema";

export const noAuthFetcher = Fetcher.for<paths>();

noAuthFetcher.configure({
  baseUrl: env.BASEURL,
  init: {},
  use: [], // middlewares
});
