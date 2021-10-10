import React, {Fragment} from 'react';
import "../../page/app.css"

export class NavBar extends React.Component {
    state = {
        mode:0,
        rank:0,
        search:'',
    }

    render() {

        return (
            <Fragment>
                <div className={"nav-bar-main"}>
                    <div className={"align-justify nav-bar-main-se"}>
                        <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" fill="currentColor"
                             className="bi bi-search svg-m05" viewBox="0 0 16 16">
                            <path
                                d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001c.03.04.062.078.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1.007 1.007 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0z"/>
                        </svg>


                        <input className={"search-input " + (this.props.width < 1024 ? "search-input-m" : "")}
                               type={"text"} placeholder={"search any"}
                               onChange={(e) => {
                                   this.setState({search: e.target.value});
                                   this.props.onSubmit(this.state);
                               }}/>
                    </div>
                </div>
                <div className={"align-justify option-nav-main "}>
                    <div className={"option-nav align-justify"}>
                        <div>{this.state.search}</div>
                        <div>-----</div>
                        <div>2</div>

                    </div>
                </div>
            </Fragment>
        );
    }
}

export default NavBar;