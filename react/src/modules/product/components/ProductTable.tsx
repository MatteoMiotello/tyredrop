import {ColumnDef} from "@tanstack/react-table";
import React from "react";
import Table from "../../../common/components-library/Table";

const ProductTable: React.FC = () => {
    const colums: ColumnDef<{ id: number, name: string }>[] = [
        {
            accessorKey: "id"
        }, {
            accessorKey: "name"
        }
    ];

    const data = [
        {
            id: 1,
            name: "ciao"
        }
    ];

    return <Table
        hideHeader={true}
        data={data}
        columns={colums}
    />;
};