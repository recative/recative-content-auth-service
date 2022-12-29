import { describe, expect, it } from "vitest";

import { RootFetcher } from "../src/fetcher/authFetcher";
import exp from "constants";

describe("permission crud", () => {
    it('create permission', async () => {
        const resp = await RootFetcher.path('/admin/permission').method('post').create()({
           id: "test_permission",
           comment: "a permission for test",
       })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    });

    it('get permission info', async () => {
        const resp = await RootFetcher.path('/admin/permission/{permission_id}').method('get').create()({
            permission_id: "test_permission",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    });

    it('modify permission info', async () => {
        let resp = await RootFetcher.path('/admin/permission/{permission_id}').method('put').create()({
            permission_id:"test_permission",
            id: "test_permission",
            comment: "after modify",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()

        resp = await RootFetcher.path('/admin/permission/{permission_id}').method('get').create()({
            permission_id: "test_permission",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    });

    it('delete permission', async () => {
        const resp = await RootFetcher.path('/admin/permission/{permission_id}').method('delete').create()({
            permission_id: "test_permission",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    })

    it('get all permissions', async () => {
        const resp = await RootFetcher.path('/admin/permissions').method('get').create()({
            ids: "test_permission_1,test_permission_2",
            regex: "test_permission_.*",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    })
});