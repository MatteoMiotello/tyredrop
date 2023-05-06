import React from "react";
import Searchbar from "./components/Searchbar";

const ProductTemplate: React.FC = () => {
    return <>
        <Searchbar/>
        <div className="w-full m-0 lg:p-24 p-4 h-full w-full flex flex-col">
        </div>;
    </>;
};

export default ProductTemplate;