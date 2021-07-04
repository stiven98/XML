import React from "react";
import ProductsApi from "../api/ProductsApi";
import ProductCard from "./ProductCard";

class Home extends React.Component {

    constructor(props) {
        super(props);
        this.state = {products: []};
    }

    componentDidMount = async () => {
        if (localStorage.getItem('id') == null) {
            this.props.history.push('/forbidden');
            return;
        }

        await ProductsApi.get('/all').then(response => {
            console.log(response);
            this.setState({products: response.data});
        })

    }

    renderProductsCards = () => {
        return this.state.products.map(product => {
            if (product.deleted === false) {
                return (<ProductCard key={product.id} product={product}/>);
            }
            return <div/>;
        })
    }

    render() {
        return (
            <div className={`m-4`}>
                <div className={`d-flex justify-content-center`}>
                    <h2 className={`text-dark font-weight-bold`}>Products:</h2>
                </div>
                <div className={`row`}>
                    {this.renderProductsCards()}
                </div>
            </div>
        )
    }

}

export default Home;
