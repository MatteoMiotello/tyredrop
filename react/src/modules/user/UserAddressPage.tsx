import {faPencil, faPlus, faTimes} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import React, { useState} from "react";
import {useTranslation} from "react-i18next";
import {useSelector} from "react-redux";
import {UserAddress} from "../../__generated__/graphql";
import Button from "../../common/components-library/Button";
import Table from "../../common/components-library/Table";
import useModal from "../../hooks/useModal";
import UserAddressModal from "./components/UserAddressModal";
import userSelectors from "./store/user-selector";

type UserAddressRowData = UserAddress
const UserAddressPage: React.FC = () => {
    const userAddresses = useSelector(userSelectors.addresses);
    const {t} = useTranslation();
    const [address, setAddress] = useState<UserAddress | undefined>(undefined);

    const {openModal, closeModal} = useModal( <UserAddressModal closeModal={() => closeModal()} address={address}/> );

    const columns: ColumnDef<UserAddressRowData>[] = [
        {
            accessorKey: 'addressName',
            header: t('user_address.user_address_name_column') as string
        },
        {
            accessorKey: 'addressLine1',
            header: t('user_address.user_address_line1_column') as string
        },
        {
            accessorKey: 'addressLine2',
            header: t('user_address.user_address_line2_column') as string
        },
        {
            accessorKey: 'city',
            header: t('user_address.user_address_city_column') as string
        },
        {
            accessorKey: 'postalCode',
            header: t('user_address.user_address_postal_code_column') as string
        },
        {
            accessorKey: 'province',
            header: t('user_address.user_address_province_column') as string
        },
        {
            accessorKey: 'buttons',
            header: ' ',
            cell: (props: CellContext<UserAddress, any>) => {
                return <>
                    <Button className="ml-auto mr-2"
                            onClick={() => {
                                openModal( <UserAddressModal closeModal={() => closeModal()} address={props.row.original}/> );
                            }}>
                        <FontAwesomeIcon icon={faPencil}/>
                    </Button>
                    <Button className="ml-auto" type="error">
                        <FontAwesomeIcon icon={faTimes}/>
                    </Button>
                </>;
            }
        }
    ];

    return <main className="p-4">
        <div className="flex w-full justify-between">
            <h3 className="text-xl">{t('user_address.title_page')}</h3>
            <Button onClick={() => {
                openModal();
            }} type="primary"> <FontAwesomeIcon icon={faPlus}/> </Button>
        </div>
        <Table data={userAddresses} columns={columns} hidePagination={true}/>
    </main>;
};

export default UserAddressPage;