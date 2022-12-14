import { noAuthFetcher } from "../fetcher/noAuthFetcher";

export async function setup() {
  const sudoToken = (
    await noAuthFetcher.path("/admin/sudo").method("post").create()({
      root_token: process.env.ROOT_TOKEN!,
    })
  ).data.token;

  process.env.SUDO_TOKEN = sudoToken;
}
