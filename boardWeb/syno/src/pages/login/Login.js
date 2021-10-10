import React from "react";
import "./Login.css"

export default class Login extends React.Component {
    render() {
        return (
            <div className={"main"}>
                <div className={"header"}>HEADER</div>
                <div className={"loginBox"}>
                    <div>
                        <div className={"tab-wrapper"} onClick={{

                        }}>
                            LOGIN
                            LOGIN
                        </div>
                    </div>
                    <div>ID</div>
                    <div>PW</div>
                </div>
                <div className={"footer"}>FOOTER</div>
            </div>
        )
    }

}