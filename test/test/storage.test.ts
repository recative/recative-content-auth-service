import { describe, expect, it } from "vitest";
import {RootFetcher} from "../src/fetcher/authFetcher";

describe("storage crud", () => {
    it('create storage', async () => {
        const resp = await RootFetcher.path('/admin/storage').method('post').create()({
            key:"114514",
            value:"1919810",
            need_permission_count:1,
            need_permissions: ["test_permission_1"],
            comment:"???",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    });

    it('get storage', async () => {
        const resp = await RootFetcher.path('/admin/storage/{storage_key}').method('get').create()({
            storage_key:"%40N3ERMcqZXqWYgrMCDj0lq%2F237%2Fabstract",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    });

    it('update storage', async () => {
        const resp = await RootFetcher.path('/admin/storage/{storage_key}').method('put').create()({
            storage_key:"114514",
            key:"114514",
            value:"114514",
            need_permission_count:2,
            need_permissions: ["test_permission_1","test_permission_2"],
            comment:"???",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()

        const resp1 = await RootFetcher.path('/admin/storage/{storage_key}').method('get').create()({
            storage_key:"114514",
        })

        expect (resp1.ok).toBe(true)
        expect (resp1.data).matchSnapshot()
    });

    it('delete storage', async () => {
        const resp = await RootFetcher.path('/admin/storage/{storage_key}').method('delete').create()({
            storage_key:"114514",
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    });

    it('get storages by query', async () => {
        const resp = await RootFetcher.path('/admin/storages').method('get').create()({
            include_value:true,
            exclude_permission:"",
            include_permission:"test_permission_1",
            keys:"a,aa"
        })

        expect (resp.ok).toBe(true)
        expect (resp.data).matchSnapshot()
    })
});