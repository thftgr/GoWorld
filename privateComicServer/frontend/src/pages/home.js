import React, {useState} from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import {Col, Container, Image, Row} from "react-bootstrap";
import {BsFillFolderFill} from "react-icons/bs";
import Modal from 'react-bootstrap/Modal'
import queryString from "query-string";
// import { useLocation, useParams } from 'react-router';
import {    BrowserRouter as Router,    Link,    useLocation} from "react-router-dom";

export default class Home extends React.Component {
    data = {}
    state = {modalImage: []}


    componentDidMount() {
        const { search } = window.location
        const {path = "/"} = queryString.parse(search)
        console.log(path )
        fetch(`http://192.168.0.62/api/file/list?path=${path }`).then((res) => res.json()).then((res) => {
            this.data = res
            this.forceUpdate()
            console.log("CALL API")
        })
    }

    galleryPreview = (props) => {
        const [fullscreen, setFullscreen] = useState(true);
        const [show, setShow] = useState(false);

        function handleShow(breakpoint) {
            setFullscreen(breakpoint);
            setShow(true);
        }

        const {thumb, galleryName} = props.gallery

        const loadImage = async () => {
            this.setState({modalImage: []})
            fetch(`http://192.168.0.62/api/file/list?path=${this.data.path}/${galleryName}`).then(res => res.json()).then(res => {
                this.setState({modalImage: res.files})
            })
        }

        return (<>
                <Image
                    src={`http://192.168.0.62/api/file/${this.data.path}/${thumb}`}
                    rounded={true}
                    style={{
                        width: '100%', maxWidth: '300px', height: '300px', objectFit: 'cover',
                    }}
                    loading="lazy"
                    onClick={() => {
                        handleShow(true)
                        loadImage()
                    }

                    }
                />
                <Modal show={show} fullscreen={fullscreen} onHide={() => setShow(false)}>
                    <Modal.Header closeButton>
                        <Modal.Title>{galleryName}</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        {this.state.modalImage.map((v, i) => {
                            return (<Image
                                    key={i}
                                    src={`http://192.168.0.62/api/file/${this.data.path}/${galleryName}/${v}`}
                                    rounded={true}
                                    style={{
                                        width: '100%', objectFit: 'cover',
                                    }}
                                    loading="lazy"/>)
                        })}
                    </Modal.Body>
                </Modal>
            </>);
    }


    render() {
        console.log(this.state)
        console.log(this.data.path)
        console.log(window.location)

        return (<Container
                style={{
                    padding: '10px', maxWidth: 1920, background: "gray", margin: `0 auto`
                }}
                fluid={true}
            >
                <Row
                    lg={"auto"}
                    md={"auto"}
                    sm={"auto"}
                    xl={"auto"}
                    // xs={"auto"}
                    xxl={"auto"}
                    // lg={{cols: 4}}
                    // md={{cols: 3}}
                    // sm={{cols: 2}}
                    // xl={{cols: 4}}
                    xs={{cols: 1}}
                    // xxl={{cols: 5}}
                    style={{
                        '--bs-gutter-x': `0em`, margin: `0 auto`, justifyContent: 'center', alignItems: 'center',

                    }}

                >
                    {this.data?.dirs?.map((v, i) => {
                        return (
                            <a href={`?path=${this.data.path === "/" ? "/": `${this.data.path}/`}${v}`}>
                                <Col key={v} style={{
                                    padding: '10px', background: "blue", // maxHeight:'300px',
                                }}>
                                    <BsFillFolderFill style={{width: '100%', height: '300px'}}/>
                                    <h3>{v}</h3>
                                </Col>
                            </a>

                            )
                    })}
                    {this.data?.galleries?.map((v, i) => {

                        return (<Col key={i} style={{
                                padding: '10px', background: "blue", // maxHeight:'300px',
                            }}>

                                <this.galleryPreview gallery={v}/>
                                <h3>{v.galleryName}</h3>
                            </Col>)
                    })}


                </Row>

            </Container>);
    }


};
