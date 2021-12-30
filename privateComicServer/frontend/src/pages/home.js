import React, {useState} from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import {Button, Col, Container, Image, Row} from "react-bootstrap";
import {BsFillFolderFill} from "react-icons/bs";
import Modal from 'react-bootstrap/Modal'

export default class Home extends React.Component {
    data = {
        // "path": "/",
        // "files": [""],
        // "galleries": [
        //     {
        //         "thumb": "12313 aaaa/bg(CUnet)(noise_scale)(Level3)(tta)(3840x2160).png",
        //         "galleryName": "12313 aaaa"
        //     },
        //     {
        //         "thumb": "12313 zzzz/bg(CUnet)(noise_scale)(Level3)(tta)(3840x2160).png",
        //         "galleryName": "12313 zzzz"
        //     }
        // ],
        // "dirs": [
        //     "12313 zzzz"
        // ]
    }
    state ={
        modalImage:[]
    }
    componentDidMount() {
        fetch("http://localhost/api/file/list").then((res) => res.json()).then((res) => {
            this.data = res
            this.forceUpdate()
            console.log("CALL API")
        })
    }
    galleryPreview = (props) => {
        const values = [true, 'sm-down', 'md-down', 'lg-down', 'xl-down', 'xxl-down'];
        const [fullscreen, setFullscreen] = useState(true);
        const [show, setShow] = useState(false);
        function handleShow(breakpoint) {
            setFullscreen(breakpoint);
            setShow(true);
        }
        const { thumb,galleryName } = props.gallery

        const loadImage = async () => {
            this.setState({modalImage: []})
            fetch(`http://localhost/api/file/list?path=${this.data.path}/${galleryName}`).then(res => res.json()).then(res => {
                this.setState({modalImage: res.files})
            })
        }

        return (
            <>
                <Image
                    src={`http://localhost/api/file/${this.data.path}/${thumb}`}
                    rounded={true}
                    style={{
                        width: '100%',
                        maxWidth: '300px',
                        height: '300px',
                        objectFit: 'cover',
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
                        {this.state.modalImage.map((v,i)=>{
                           return(
                               <Image
                                   key={i}
                                   src={`http://localhost/api/file/${this.data.path}/${galleryName}/${v}`}
                                   rounded={true}
                                   style={{
                                       width: '100%',
                                       objectFit: 'cover',
                                   }}
                                   loading="lazy"/>
                           )
                        })}
                    </Modal.Body>
                </Modal>
            </>
        );
    }

    render() {
        console.log(this.state)
        console.log(this.data.path)

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
                                    // window.history.push(`/?path=path=${this.data.path}/${v}`)

                                    fetch(`http://localhost/api/file/list?path=${this.data.path}/${v}`).then((res) => res.json()).then((res) => {
                                        this.data = res
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
                        this.data?.galleries?.map((v, i) => {

                            return (
                                <Col key={i} style={{
                                    padding: '10px',
                                    background: "blue",
                                    // maxHeight:'300px',
                                }}>
                                    {/*<Image*/}
                                    {/*    src={`http://localhost/api/file${this.data.path}/${v.thumb}`}*/}
                                    {/*    rounded={true}*/}
                                    {/*    style={{*/}
                                    {/*        width: '100%',*/}
                                    {/*        maxWidth: '300px',*/}
                                    {/*        height: '300px',*/}
                                    {/*        objectFit: 'cover',*/}
                                    {/*    }}*/}
                                    {/*    loading="lazy"*/}
                                    {/*    // onClick={}*/}
                                    {/*/>*/}

                                    <this.galleryPreview gallery={v} />
                                    {/*<BsImage style={{width:'100%',height:'300px'}}/>*/}
                                    <h3>gal: {v.galleryName}</h3>
                                </Col>
                            )
                        })
                    }


                </Row>

            </Container>
        );
    }


};
