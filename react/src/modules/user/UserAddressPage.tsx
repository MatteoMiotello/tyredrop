import {faPencil, faPlus, faTimes} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import React from "react";
import {useTranslation} from "react-i18next";
import {useDispatch, useSelector} from "react-redux";
import {UserAddress} from "../../__generated__/graphql";
import Button from "../../common/components-library/Button";
import Table from "../../common/components-library/Table";
import useModal from "../../hooks/useModal";
import {useToast} from "../../hooks/useToast";
import ConfirmDeleteModal from "./components/ConfirmDeleteModal";
import UserAddressModal from "./components/UserAddressModal";
import userSelectors from "./store/user-selector";
import {deleteUserAddress} from "./store/user-slice";
import {ThunkDispatch} from "redux-thunk";

type UserAddressRowData = UserAddress
const UserAddressPage: React.FC = () => {
    const userAddresses = useSelector(userSelectors.addresses);
    const {t} = useTranslation();
    const {setSuccess, setError} = useToast();
    const {openModal, closeModal} = useModal();
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();

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
            accessorKey: 'country',
            header: t('user_address.user_address_country_column') as string
        },
        {
            accessorKey: 'buttons',
            header: ' ',
            cell: (props: CellContext<UserAddress, any>) => {
                return <div className="flex">
                    <Button className="mx-1"
                            onClick={() => {
                                openModal(<UserAddressModal closeModal={() => closeModal()}
                                                            address={props.row.original}/>);
                            }}>
                        <FontAwesomeIcon icon={faPencil}/>
                    </Button>
                    <Button className="mx-1" type="error" onClick={() => {
                        openModal(<ConfirmDeleteModal
                            closeModal={closeModal}
                            onConfirm={
                                () => {
                                    dispatch(deleteUserAddress({id: props.row.original.ID}))
                                        .unwrap()
                                        .then( () => setSuccess( t( 'user_address.delete_success' ) ) )
                                        .catch( () => setError( t( 'user_address.delete_error' ) ) );
                                    closeModal();
                                }

                        }/>
                        );
                    }}>
                        <FontAwesomeIcon icon={faTimes}/>
                    </Button>
                </div>;
            }
        }
    ];

    return <main className="p-4">
        <div className="flex w-full justify-between">
            <h3 className="text-xl">{t('user_address.title_page')}</h3>
            <Button onClick={() => {
                openModal(<UserAddressModal closeModal={() => closeModal()}/>);
            }} type="primary"> <FontAwesomeIcon icon={faPlus}/> </Button>
        </div>
        <Table data={userAddresses} columns={columns} hidePagination={true}/>
    </main>;
};

export default UserAddressPage;