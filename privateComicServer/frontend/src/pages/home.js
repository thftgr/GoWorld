import React, {useState} from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import {Button, Col, Container, Image, Row} from "react-bootstrap";
import {BsFillFolderFill} from "react-icons/bs";
import Modal from 'react-bootstrap/Modal'

export default class Home extends React.Component {
    state = {
        baseUrl: []
    }
    data = {
        // "dirs": [
        //     "12313 zzzz"
        // ],
        // "gallery": [
        //     {
        //         "Name": "12313 aaaa",
        //         "Thumbnail": "/12313 aaaa/bg(CUnet)(noise_scale)(Level3)(tta)(3840x2160).png"
        //     },
        //     {
        //         "Name": "12313 zzzz",
        //         "Thumbnail": "/12313 zzzz/bg(CUnet)(noise_scale)(Level3)(tta)(3840x2160).png"
        //     }
        // ],
        // "path": ""
    }

    componentDidMount() {
        fetch("http://localhost/api/file/list").then((res) => res.json()).then((res) => {
            this.data = res
            this.forceUpdate()
            console.log("CALL API")
        })
    }

    Example = () => {
        const values = [true, 'sm-down', 'md-down', 'lg-down', 'xl-down', 'xxl-down'];
        const [fullscreen, setFullscreen] = useState(true);
        const [show, setShow] = useState(false);

        function handleShow(breakpoint) {
            setFullscreen(breakpoint);
            setShow(true);
        }

        return (
            <>
                {values.map((v, idx) => (
                    <Button key={idx} className="me-2" onClick={() => handleShow(v)}>
                        Full screen
                        {typeof v === 'string' && `below ${v.split('-')[0]}`}
                    </Button>
                ))}
                <Modal show={show} fullscreen={fullscreen} onHide={() => setShow(false)}>
                    <Modal.Header closeButton>
                        <Modal.Title>Modal</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>Modal body content</Modal.Body>
                </Modal>
            </>
        );
    }


    render() {
        console.log(this.state)
        return (
            <Container
                style={{
                    padding: '10px',
                    maxWidth: 1920,
                    background: "gray",
                    margin: `0 auto`
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
                        '--bs-gutter-x': `0em`,
                        margin: `0 auto`,
                        justifyContent: 'center',
                        alignItems: 'center',

                    }}

                >
                    {
                        this.data?.dirs?.map((v, i) => {
                            return (
                                <Col key={v} style={{
                                    padding: '10px',
                                    background: "blue",
                                    // maxHeight:'300px',
                                }} onClick={() => {
                                    fetch(`http://localhost/api/file/list?path=${this.data.path}/${v}`).then((res) => res.json()).then((res) => {
                                        this.data = res
                                        this.forceUpdate()
                                        this.forceUpdate()
                                    })
                                }}>
                                    <BsFillFolderFill style={{width: '100%', height: '300px'}}/>
                                    <h3>dir: {v}</h3>
                                </Col>
                            )
                        })
                    }
                    {
                        this.data?.gallery?.map((v, i) => {
                            return (
                                <Col key={i} style={{
                                    padding: '10px',
                                    background: "blue",
                                    // maxHeight:'300px',
                                }}>
                                    <Image
                                        src={`http://localhost/api/file${v.Thumbnail}`}
                                        rounded={true}
                                        style={{
                                            width: '100%',
                                            maxWidth: '300px',
                                            height: '300px',
                                            objectFit: 'cover',
                                        }}
                                        loading="lazy"
                                        onClick={() => {
                                            const [fullscreen] = useState(true);
                                            const [show, setShow] = useState(false);
                                            return(
                                                <Modal show={show} fullscreen={fullscreen} onHide={() => setShow(false)}>
                                                    <Modal.Header closeButton>
                                                        <Modal.Title>Modal</Modal.Title>
                                                    </Modal.Header>
                                                    <Modal.Body>Modal body content</Modal.Body>
                                                </Modal>
                                            )

                                        }}
                                    />
                                    {/*<BsImage style={{width:'100%',height:'300px'}}/>*/}
                                    <h3>gal: {v.Name}</h3>
                                </Col>
                            )
                        })
                    }


                </Row>

            </Container>
        );
    }


};
