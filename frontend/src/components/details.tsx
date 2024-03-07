import QRCode from "qrcode";
import React, { useRef, useEffect, useCallback } from "react";
import * as types from "../../@types";

export function Details({ item }: { item: types.GetShortcodesResponse; }): JSX.Element | null {

    const modalRef = useRef<HTMLDialogElement | null>(null)
    const canvasRef = useRef<HTMLCanvasElement | null>(null)

    useEffect(() => {
        (async () => {
            QRCode.toCanvas(canvasRef.current, item.shortUrl, function (error) {
                if (error) {
                    console.error(error)
                } else {
                    console.log("success!");
                }
            })
        })()
    },[item.shortUrl])

    const generateAsciiQrCode = useCallback(() => {
        QRCode.toString(item.shortUrl, { type: "terminal" }, function (error, url) {
            if (error) {
                console.error(error)
            } else {
                console.log(url)
            }
        })
    },[item.shortUrl])
    
    return (
        <>
            <dialog ref={modalRef} className="modal">
                <div className="modal-box">
                    <h3 className="font-bold text-lg">Details</h3>
                    <p className="py-4">Original Url</p>
                    <p className="py-4">{item.originalUrl}</p>
                    <p className="py-4">Short Url</p>
                    <p className="py-4">{item.shortUrl}</p>
                    <p className="py-4">QR Code</p>
                    <canvas ref={canvasRef} />

                    <p className="py-4">Ascii QR Code</p>
                    <button className="btn" onClick={generateAsciiQrCode}>Generate in Console</button>            

                    <div className="modal-action">
                        <button className="btn" onClick={() => modalRef.current?.close("close")}>Close</button>
                    </div>
                </div>
            </dialog>
            <button onClick={() => modalRef.current?.showModal()} className="btn btn-ghost btn-xs">details</button>
        </>
    )
}