import React, { ReactElement } from "react";

interface ButtonProps {
    onClick() : void
    text? : ReactElement
    className? : string
}

function Button(props : ButtonProps) {
    return (
        <button onClick={props.onClick} className={props.className}>
            {props.text}
        </button>
    )
}

export default Button;