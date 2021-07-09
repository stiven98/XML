import React from "react";
import {CampaignModel} from "../model/CampaignModel";
import axios from "axios";
import AdModel from "../model/AdModel"
import NistagramApi from "../api/NistagramApi";
import { withRouter } from "react-router-dom";

class Campaign extends React.Component {

    constructor(props) {
        super(props);
        this.state = {campaign: new CampaignModel()};
    }

    changeMultiple = (event) => {
        this.setState({campaign: {...this.state.campaign, ismultiple: !this.state.campaign.ismultiple}});
    }

    changeMen = () => {
        this.setState({campaign: {...this.state.campaign, showtomen: !this.state.campaign.showtomen}});
    }

    changeWomen = () => {
        this.setState({campaign: {...this.state.campaign, showtowomen: !this.state.campaign.showtowomen}});
    }

    changeUnder18 = () => {
        this.setState({campaign: {...this.state.campaign, showunder18: !this.state.campaign.showunder18}});
    }

    change18to24 = () => {
        this.setState({campaign: {...this.state.campaign, show18to24: !this.state.campaign.show18to24}});
    }

    changeOver25 = () => {
        this.setState({campaign: {...this.state.campaign, showover35: !this.state.campaign.showover35}});
    }

    change24to35 = () => {
        this.setState({campaign: {...this.state.campaign, show24to35: !this.state.campaign.showover24to35}});
    }

    onChangeStartTime = (event) => {
        this.setState({campaign: {...this.state.campaign, startday: event.target.value}});
    }

    onChangeEndTime = (event) => {
        this.setState({campaign: {...this.state.campaign, endday: event.target.value}});
    }

    renderDates = () => {
        if (this.state.campaign.ismultiple === true) {
            return (
                <div>
                    <div className="input-group mb-3 px-2 py-2 rounded-pill bg-white shadow-sm">
                        <label htmlFor="location" className="font-weight-light text-muted mt-2 ml-3">Datum
                            početka:</label>
                        <input type="date" id="location" className="form-control border-0" name="location"
                               onChange={this.onChangeStartTime} value={this.state.campaign.startday}/>
                    </div>
                    <div className="input-group mb-3 px-2 py-2 rounded-pill bg-white shadow-sm">
                        <label htmlFor="location" className="font-weight-light text-muted mt-2 ml-3">Datum
                            isticanja:</label>
                        <input type="date" id="location" className="form-control border-0" name="location"
                               onChange={this.onChangeEndTime} value={this.state.campaign.endday}/>
                    </div>
                </div>
            )
        }
    }

    onChangeDescription = (event) => {
        this.setState({campaign: {...this.state.campaign, description: event.target.value}})
    }

    onChangeTimesToPlace = (event) => {
        this.setState({campaign: {...this.state.campaign, timestoplace: event.target.value}})
    }

    onChangeWhenToPlace = (event) => {
        this.setState({campaign: {...this.state.campaign, whentoplace: event.target.value}})
    }

    upload = (event) => {
        event.preventDefault();
        const formData = new FormData();
        formData.append('files', event.target.files[0]);
        NistagramApi.post('upload', formData).then(response => {
            const ads = this.state.campaign.ads;
            let newAds = new AdModel();
            newAds.path = response.data[0].path;
            ads.push(newAds);
            this.setState({campaign: {...this.state.campaign, ads: ads}})
        });
    }

    submit = async (event) => {
        event.preventDefault();

        await axios.get('http://localhost:8085/sysusers/getUserId/' + localStorage.getItem('username')).then(response => {
            this.setState({campaign: {...this.state.campaign, userId: response.data}});
        });

        await this.setState({
            campaign: {
                ...this.state.campaign,
                startday: this.state.campaign.startday + 'T01:00:00+01:00'
            }
        });
        await this.setState({
            campaign: {
                ...this.state.campaign,
                endday: this.state.campaign.endday + 'T01:00:00+01:00'
            }
        });




        if (this.state.campaign.ads.length === 1) {
            this.setState({campaign: {...this.state.campaign, type: 'post'}});

        } else if (this.state.ads.length > 1) {
            this.setState({campaign: {...this.state.campaign, type: 'album'}});
        }

        console.log(this.state.campaign);


        await NistagramApi.post('campaigns/createCampaign', this.state.campaign).then((response => {
            this.setState({campaign: new CampaignModel()});
            alert("Success!");
            window.location.href.push('/');
        })).catch(err => {
            console.log(err);
        })
    }

    onChangeLink = (event, index) => {
        const ads = this.state.campaign.ads;
        ads[index].link = event.target.value;
        this.setState({campaign: {...this.state.campaign, ads: ads}});
    }

    renderAds = () => {
        return this.state.campaign.ads.map((item, index) => {
            return (
                <div key={index}>
                    <label htmlFor="hashtag" className={`mr-3`}>{item.path}</label>
                    <input type="text" className="form-control border-0" value={item.link}
                           onChange={(event) => this.onChangeLink(event, index)}/>
                </div>


            );
        })
    }

    render() {
        return (
            <div>
                <div className={`d-flex justify-content-center mt-3`}>
                    <h1 className={`text-dark`}>Campaign:</h1>
                </div>

                <div className={`d-flex justify-content-center mt-4`}>
                    <form className={`form-control-feedback w-50`}>
                        <div className=" d-flex justify-content-center mt-3 mb-3    ">
                            <label htmlFor="hashtag" className={`mr-3`}>Da li je kampanja
                                višekratna?</label>
                            <label className="form-check-label mr-2" htmlFor="isMultiple">Da</label>
                            <input type="checkbox" id="isMultiple" className={`mb-1`} onChange={this.changeMultiple}/>
                        </div>

                        {this.renderDates()}


                        <div className="input-group mb-3 px-2 py-2 rounded-pill bg-white shadow-sm">
                            <label htmlFor="description"
                                   className="font-weight-light text-muted mt-2 ml-3">Opis:</label>
                            <input type="text" id="description" className="form-control border-0" name="description"
                                   value={this.state.campaign.description} onChange={this.onChangeDescription}/>
                        </div>
                        <div className="input-group mb-3 px-2 py-2 rounded-pill bg-white shadow-sm">
                            <label htmlFor="description" className="font-weight-light text-muted mt-2 ml-3">Kada
                                plasirati
                                kampanju:</label>
                            <input type="time" id="time" className="form-control border-0" name="time"
                                   value={this.state.campaign.whentoplace} onInput={this.onChangeWhenToPlace}/>

                        </div>
                        <div className="input-group mb-3 px-2 py-2 rounded-pill bg-white shadow-sm">
                            <label htmlFor="description" className="font-weight-light text-muted mt-2 ml-3">Koliko puta
                                plasirati:</label>
                            <input type="number" id="description" className="form-control border-0" name="description"
                                   value={this.state.campaign.timestoplace} onChange={this.onChangeTimesToPlace}/>
                        </div>


                        <div className="form-check  px-2 py-2 pill bg-white shadow-sm">
                            <div className={`m-4 `}>

                                <label htmlFor="description" className="text-muted">Ciljna grupa:</label>

                                <div className="form-check">
                                    <label className="form-check-label" htmlFor={`males`}>
                                        <input type="checkbox" className="form-check-input" id={`males`} value=""
                                               onChange={this.changeMen}/>Muskarci
                                    </label>
                                </div>

                                <div className="form-check">
                                    <label className="form-check-label" htmlFor={`women`}>
                                        <input type="checkbox" className="form-check-input" id={`women`} value=""
                                               onChange={this.changeWomen}/>Zene
                                    </label>
                                </div>

                                <div className="form-check">
                                    <label className="form-check-label">
                                        <input type="checkbox" className="form-check-input" value=""
                                               onChange={this.changeUnder18}/>Mladji od 18
                                        godina
                                    </label>
                                </div>

                                <div className="form-check">
                                    <label className="form-check-label">
                                        <input type="checkbox" className="form-check-input" onChange={this.change18to24}
                                               value=""/>Između 18 i 24
                                        godine
                                    </label>
                                </div>

                                <div className="form-check">
                                    <label className="form-check-label">
                                        <input type="checkbox" className="form-check-input" onChange={this.change24to35}
                                               value=""/>Između 24 i 35
                                        godine
                                    </label>
                                </div>

                                <div className="form-check">
                                    <label className="form-check-label">
                                        <input type="checkbox" className="form-check-input" value=""
                                               onChange={this.changeOver25}/>Preko 25
                                        godina
                                    </label>
                                </div>
                            </div>
                        </div>

                        <div className="input-group mb-3 px-2 py-2 rounded-pill bg-white shadow-sm">
                            <input id="upload" className="form-control border-0" type="file" name="files" multiple
                                   onChange={this.upload}/>
                            <div className="input-group-append">
                                <label htmlFor="upload" className="btn btn-light m-0 rounded-pill px-4"> <i
                                    className="fa fa-cloud-upload mr-2 text-muted"/><small
                                    className="text-uppercase font-weight-bold text-muted">Odaberite
                                    fajl</small></label>
                            </div>
                        </div>
                        <div>
                            {this.renderAds()}
                        </div>
                        <div className="d-flex justify-content-center">
                            <button className="btn btn-outline-dark" onClick={this.submit}>Objavi</button>
                        </div>
                    </form>
                </div>
            </div>
        );
    }
}

export default withRouter(Campaign);
