import { describe, expect, it } from "vitest";

import { RootFetcher } from "../src/fetcher/authFetcher";

describe("admin token crud", () => {
    const payload = {
        token: "38a44a3f-c108-412b-a910-9e47adc17c27",
    }

    it('create admin token', async () => {
        let date = new Date()
        // add 5 days
        date.setDate(date.getDate() + 5)
        const resp = await RootFetcher.path('/admin/token').method('post').create()({
            is_valid:true,
            comment:"a token for test",
            admin_permission: ["write"],
            token: payload.token,
            expired_at: date.toISOString(),
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    });

    it('get admin token info', async () => {
        const resp = await RootFetcher.path('/admin/token/{token}').method('get').create()({
            token: payload.token,
        })

        expect (resp.ok).toBe(true)
        resp.data.expired_at = undefined
        expect (resp.data).matchSnapshot()
    });

    it('update admin token', async () => {
        const resp = await RootFetcher.path('/admin/token/{token}').method('put').create()({
            token: payload.token,
            is_valid: false,
            comment: "a token for test",
            admin_permission: ["read"],
            expired_at: undefined,
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()

        const resp2 = await RootFetcher.path('/admin/token/{token}').method('get').create()({
            token: payload.token,
        })

        expect (resp2.ok).toBe(true)
        resp2.data.expired_at = undefined
        expect (resp2.data).matchSnapshot()
    })

    it('delete token', async () => {
        const resp = await RootFetcher.path('/admin/token/{token}').method('delete').create()({
            token: payload.token,
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    });

    it('get all tokens', async () => {
        const resp = await RootFetcher.path('/admin/tokens').method('get').create()({})

        expect (resp.ok).toBe(true)
        const data = resp.data.map((item) => {
            item.expired_at = undefined
            return item
        })
        expect (data).matchSnapshot()
    })
});