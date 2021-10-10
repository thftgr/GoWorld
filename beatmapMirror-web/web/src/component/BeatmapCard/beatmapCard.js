import React from 'react';
import "../component.css"
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'
import {faCoffee, faHeart, faPlayCircle, faMusic, faArchive} from '@fortawesome/free-solid-svg-icons'
import {OverlayTrigger, Tooltip} from "react-bootstrap";

// class ToolTip extends React.Component{
//     render() {
//         return(
//             <OverlayTrigger
//                 key={placement}
//                 placement={placement}
//                 overlay={
//                     <Tooltip id={`tooltip-bottom`}>
//                         Tooltip on <strong>{placement}</strong>.
//                     </Tooltip>
//                 }
//             >
//                 <Button variant="secondary">Tooltip on {placement}</Button>
//             </OverlayTrigger>
//         )
//     }
// }





export default (props) => {
    return (
        <div className={"mapSetCard"}>
            <div className={"mapSetCard-top"} style={{
                background: `url('https://assets.ppy.sh/beatmaps/` + props.setData.SetID + `/covers/cover@2x.jpg') no-repeat`,
                backgroundSize: "cover",

            }}>
                <div>
                    <div className="beatmapset-status-container">
                        <div className="beatmapset-status">{props.setData.RankedStatus_Text}</div>
                        <div className="beatmapset-pv-play">
                        <span className="beatmapset-panel-prev beatmapset-panel__play">
                            <svg xmlns="http://www.w3.org/2000/svg"
                                 fill="currentColor"
                                 className="bi bi-play-fill" viewBox="0 0 16 16">
                                <path
                                    d="M11.596 8.697l-6.363 3.692c-.54.313-1.233-.066-1.233-.697V4.308c0-.63.692-1.01 1.233-.696l6.363 3.692a.802.802 0 0 1 0 1.393z"/>
                            </svg>
                        </span>
                        </div>
                    </div>
                    <div className="beatmapset-info">
                        <Tooltip id={`tooltip-bottom`}>
                            <div className="beatmapset-info-row"><FontAwesomeIcon
                                icon={faHeart}/>{props.setData.Favourites}</div>
                        </Tooltip>
                        <Tooltip id={`tooltip-bottom`}>
                            <div className="beatmapset-info-row"><FontAwesomeIcon
                                icon={faPlayCircle}/>{props.setData.Playcounts}</div>
                        </Tooltip>
                        <Tooltip id={`tooltip-bottom`}>
                            <div className="beatmapset-info-row"><FontAwesomeIcon icon={faMusic}/>{props.setData.BPM}
                            </div>
                        </Tooltip>
                        <Tooltip id={`tooltip-bottom`}>
                            <div className="beatmapset-info-row"><FontAwesomeIcon
                                icon={faArchive}/>{props.setData.BeatmapCount}</div>
                        </Tooltip>
                    </div>
                </div>


            </div>
            <div className={"mapSetCard-bottom"}>
                <div>
                    <div>mapped by: {props.setData.Creator}</div>
                    <div></div>
                    {/*<div>Title: {props.setData.Title}</div>*/}
                    {/*<div>Artist: {props.setData.Artist}</div>*/}

                </div>
                <div>
                    <i></i>
                </div>

            </div>


        </div>
    );
}



