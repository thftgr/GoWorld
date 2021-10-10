import React, {Fragment} from 'react';
import {BeatmapCard, NavBar} from "./index";


export default class Home extends React.Component {

    offset = 0

    mapTmp = []
    state = {
        columnCount: {gridTemplateColumns: '1fr 1fr'},
        mapList: [],
        compo: null
    };

    componentDidMount() {
        //두번연속 로딩해야함
        const url = (amount) => {
            this.offset += amount
            return `https://nerina.wtf/api/v3/search?offset=${this.offset-amount}&mode=0&amount=${amount}&status=1&query=`
        }
        if (this.state.mapList.length === 0 ){
            fetch(url(24)).then(res => res.json()).then(res => {
                const tp = this.state.mapList
                tp.push(...res)
                this.setState({mapList: tp})
            })
            // this.amount =
            fetch(url(48)).then(res => res.json()).then(res => {
                this.mapTmp = res
            })
        }else {
            const tp = this.state.mapList
            tp.push(...this.mapTmp)
            this.setState({mapList: tp})
            fetch(url(48)).then(res => res.json()).then(res => {
                this.mapTmp = res
            })
        }


        window.addEventListener("scroll", this.infiniteScroll, true);
    }

    infiniteScroll = () => {
        let scrollHeight = Math.max(
            document.documentElement.scrollHeight,
            document.body.scrollHeight
        );
        let scrollTop = Math.max(
            document.documentElement.scrollTop,
            document.body.scrollTop
        );
        let clientHeight = document.documentElement.clientHeight;

        if (scrollTop + clientHeight >= scrollHeight) {
            this.componentDidMount();
        }
    };

    render() {
        const mtc = data => {
            return data.map((bc, i) => {
                return (<BeatmapCard setData={bc} width={""} key={i}/>)
            })
        }

        return (
            <Fragment>
                <NavBar onSubmit={(value) => {
                    this.setState(value)
                }} width={""}/>
                <div className={"mapCardArea"} style={this.state.columnCount}>
                    {mtc(this.state.mapList)}
                </div>
            </Fragment>
        );

    }


}
