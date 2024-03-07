import React from "react";

export function ErrorComponent({message}: {message: string}): JSX.Element {
    return (
        <div className="bg-error text-black">
            {message}
        </div>
    );
}
