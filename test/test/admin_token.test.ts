import { describe, expect, it } from "vitest";

import { SudoFetcher } from "../src/fetcher/authFetcher";

describe("admin token crud", () => {
    it('create admin token', async () => {
        let date = new Date()
        // add 5 days
        date.setDate(date.getDate() + 5)
        const resp = await SudoFetcher.path('/admin/token').method('post').create()({
            is_valid:true,
            comment:"a token for test",
            admin_permission: ["write"],
            token: undefined,
            expired_at: date.toISOString(),
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    });
});