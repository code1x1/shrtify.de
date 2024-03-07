import React, { useEffect, useMemo, useState } from "react"
import { useAdmin } from "src/hooks/admin";
import * as types from "../../@types";
import { Details } from "./details";

export default function ShortcodeList() {
    const {getShortcodes} = useAdmin();
    
    const [list, setList] = useState<types.GetShortcodesResponse[]>([])
    const [pages, setPages] = useState<number>(1)
    const [currentPage, setCurrentPage] = useState<number>(1)

    useEffect(() => {
        (async () => {
            const response = await getShortcodes({ page: currentPage, items: 2})
            if (Array.isArray(response.data.response)) {
                setList(response.data.response)
                setPages(response.data.meta.more.pages);
            }
        })()
    }, [currentPage, getShortcodes])


    const items = useMemo(() => list.map((item, index) => (
        <tr key={index}>
            <td>
                <div className="flex items-center gap-3">
                    <div>
                        {item.originalUrl}
                    </div>
                </div>
            </td>
            <td>
                <a href={item.shortUrl} target="_blank">
                    {item.shortUrl}
                </a>
            </td>
            <th>
                <Details item={item} />
            </th>
        </tr>
    )),[list])

    const pagination = useMemo(() => Array.from(new Array(pages)).map((_, index) => {
        const page = index + 1
        const onClick = () => setCurrentPage(page);
        return (
            <button key={index} onClick={onClick} className={"join-item btn" + (page === currentPage ? " btn-active" : "")}>{page}</button>
        )
    }),[currentPage, pages])

    return (
        <div>

            <div className="overflow-x-auto">
                <table className="table">
                    <thead>
                        <tr>
                            <th>Original Url</th>
                            <th>Short Url</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {items}
                    </tbody>
                </table>
            </div>
            <div className="flex flex-row justify-center">
                <div className="join flex-grow-0">
                    {pagination}
                </div>
            </div>
        </div>
    )
}