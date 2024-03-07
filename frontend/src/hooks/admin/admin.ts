import { useCallback } from "react";
import api from "../api/api";
import * as types from "../../../@types";

export function useAdmin() {
    const createShortcode = useCallback(
        async (request: types.CreateShortcodeRequestPayload["request"]) => {
            const payload: types.CreateShortcodeRequestPayload = {
                request,
            };
            return await api.post<
                types.CreateShortcodeRequestPayload,
                {
                    data: types.CreateShortcodeResponsePayload;
                }
            >("/api/v1/admin/createShortcode", payload);
        },
        []
    );

    const getShortcodes = useCallback(
        async (options: { page: number; items: number }) => {
            return await api.get<types.GetShortcodesResponsePayload>(
                `/api/v1/admin/getShortcodes?page=${options.page}&items=${options.items}`
            );
        },
        []
    );

    return {
        createShortcode,
        getShortcodes,
    };
}
