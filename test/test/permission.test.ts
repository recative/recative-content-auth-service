import { describe, expect, it } from "vitest";

import { SudoFetcher } from "../src/fetcher/authFetcher";

describe("permission crud", () => {
    it('create permission', async () => {
        const resp = await SudoFetcher.path('/admin/permission').method('post').create()({
           id: "test_permission",
           comment: "a permission for test",
       })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    });

    it('get permission info', async () => {
        const resp = await SudoFetcher.path('/admin/permission/{permission_id}').method('get').create()({
            permission_id: "test_permission",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    });

    it('modify permission info', async () => {
        let resp = await SudoFetcher.path('/admin/permission/{permission_id}').method('put').create()({
            permission_id:"test_permission",
            id: "test_permission",
            comment: "after modify",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()

        resp = await SudoFetcher.path('/admin/permission/{permission_id}').method('get').create()({
            permission_id: "test_permission",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    });

    it('delete permission', async () => {
        const resp = await SudoFetcher.path('/admin/permission/{permission_id}').method('delete').create()({
            permission_id: "test_permission",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    })

    it('batch get permission', async () => {
        // const resp = await SudoFetcher.path('/admin/permissions').method('post').create()({
        //
        // })
        // const resp = await SudoFetcher.path('/admin/permission/{permission_id}').method('delete').create()({
        //     permission_id: "test_permission",
        // })
        //
        // expect (resp.ok).toBe(true)
        // expect (resp.data).matchSnapshot()
    })
});