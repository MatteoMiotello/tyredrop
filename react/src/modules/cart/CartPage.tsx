import React from "react";
import CartTable from "./components/CartTable";

const CartPage: React.FC = () => {
    return <main className="w-full m-0 lg:px-24 px-4 h-full flex flex-col min-h-screen">
        <CartTable/>
    </main>;
};

export default CartPage;