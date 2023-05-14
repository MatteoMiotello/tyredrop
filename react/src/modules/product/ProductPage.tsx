import React from "react";
import Searchbar from "./components/Searchbar";
import Table from "../../common/components-library/Table";
import {ColumnDef} from "@tanstack/react-table";

const ProductPage: React.FC = () => {
    const colums: ColumnDef<{ id: number, name: string }>[] = [
        {
            header: "ID",
            accessorKey: "id"
        },
        {
            header: "Nome",
            accessorKey: "name"
        },
    ];

    const data = [
        
    ];

    return <>
        <Searchbar/>
        <div className="w-full m-0 lg:px-24 px-4 h-full w-full flex flex-col">
            <Table
                data={data}
                columns={colums}
            />
        </div>
    </>;
};

export default ProductPage;