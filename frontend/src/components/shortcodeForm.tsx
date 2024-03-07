import React, { useCallback, useMemo, useState } from "react";
import { useForm } from "react-hook-form";

import { ErrorMessage } from "@hookform/error-message";
import { ErrorComponent } from "./subtopics";
import { useAdmin } from "../hooks/admin/admin";
import * as types from "../../@types";
import ShortcodeList from "./shortcodeList";

export default function App() {
    const {
        register,
        handleSubmit,
        formState: {isSubmitSuccessful, errors},
        reset,
        resetField,
    } = useForm<types.CreateShortcodeRequest>({
        defaultValues: {
            url: "",
            code: "",
        },
    });
    const {createShortcode} = useAdmin();
    const [shortUrl, setShortUrl] = useState("")

    React.useEffect(() => {
        if (isSubmitSuccessful) {
            resetField("url", {
                defaultValue: "",
            });
            resetField("code", {
                defaultValue: "",
            });
        }
    }, [isSubmitSuccessful, reset, resetField]);

    const onSubmit = useCallback(async (data: types.CreateShortcodeRequest) => {
        const payload = await createShortcode(data);
        setShortUrl(payload.data.response.shortUrl)
    }, [createShortcode]);

    const url = useMemo(() => {
        const result = ({
            ...register("url", {
                required: {
                    value: true,
                    message: "The Url is required.",
                },
                pattern: {
                    value: /https?:\/\/(\w+)/i,
                    message: "The Url protocol is required.",
                }}),
            type: "text",
            placeholder: "Please enter target address.",
            className: "",
        });
        if (errors.url) {
            result.className += "border-error border-b-2";
        }
        return result;
    }, [errors.url, register]);

    const code = useMemo(() => {
        const result = ({
            ...register("code"),
            type: "text",
            placeholder: "Please enter link code. (Optional)",
            className: "",
        });
        if (errors.code) {
            result.className += "border-error border-b-2";
        }
        return result;
    }, [errors.code, register]);

    return (
        <div>
            <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-y-3">
                <input {...url} />
                <ErrorMessage
                    name="url"
                    errors={errors}
                    render={ErrorComponent} />
                <input {...code} />
                <ErrorMessage
                    name="code"
                    errors={errors}
                    render={ErrorComponent} />

                <input type="submit" value="Senden" className="btn" />
                {shortUrl ? <a href={shortUrl}>{shortUrl}</a> : null}
            </form>
            <ShortcodeList />
        </div>
    );
}
