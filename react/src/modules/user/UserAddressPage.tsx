import {useQuery} from "@apollo/client";
import {faPencil, faPlus, faTimes} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import React, {useState} from "react";
import {useTranslation} from "react-i18next";
import {useDispatch} from "react-redux";
import {useParams} from "react-router-dom";
import {UserAddress} from "../../__generated__/graphql";
import {USER_ADDRESSES} from "../../common/backend/graph/query/users";
import Button from "../../common/components-library/Button";
import Panel from "../../common/components-library/Panel";
import Spinner from "../../common/components/Spinner";
import ConfirmModal from "./components/ConfirmModal";
import {deleteUserAddress} from "./store/user-slice";
import {ThunkDispatch} from "redux-thunk";
import {useToast} from "../../store/toast";
import {BasicTable, useModal, useTable} from "../../common/components/shelly-ui";
import UserAddressModal from "./components/UserAddressModal";

type UserAddressRowData = UserAddress
const UserAddressPage: React.FC = () => {
    const params = useParams<{ id: string }>();
    const {data, loading, error, refetch} = useQuery(USER_ADDRESSES, {
        variables: {
            userId: params.id
        }
    });
    const {t} = useTranslation();
    const toastr = useToast();
    const userAddressModal = useModal({
        onClose: () => refetch()
    });
    const confirmDeleteModal = useModal();
    const [selectedAddress, setSelectedAddress] = useState<UserAddress | undefined>();
    const [selectedAddressToDelete, setSelectedAddressToDelete] = useState<UserAddress | undefined>();
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
                                setSelectedAddress(props.row.original);
                                userAddressModal.open();
                            }}>
                        <FontAwesomeIcon icon={faPencil}/>
                    </Button>
                    <Button className="mx-1" type="error" onClick={() => {
                        setSelectedAddressToDelete(props.row.original);
                        confirmDeleteModal.open();
                    }}>
                        <FontAwesomeIcon icon={faTimes}/>
                    </Button>
                </div>;
            }
        }
    ];

    const {table} = useTable({
        data: data?.userAddress,
        columns: columns,
    });

    if (loading) {
        return <Spinner/>;
    }

    return <main className="h-full">
        <UserAddressModal modal={userAddressModal} address={selectedAddress}/>
        <ConfirmModal modal={confirmDeleteModal} onConfirm={() => {
            if (!selectedAddressToDelete) {
                return;
            }

            dispatch(deleteUserAddress({id: selectedAddressToDelete?.ID}))
                .unwrap()
                .then(() => {
                    toastr.success(t('user_address.delete_success'));
                    refetch();
                })
                .catch(() => toastr.error(t('user_address.delete_error')));

            confirmDeleteModal.close();
        }}/>
        <Panel>
            <div className="flex w-full justify-between">
                <h3 className="text-xl">{t('user_address.title_page')}</h3>
                <Button onClick={() => {
                    setSelectedAddress(undefined);
                    userAddressModal.open();
                }} type="primary"> <FontAwesomeIcon icon={faPlus}/> </Button>
            </div>
            {
                data &&
                <BasicTable table={table}/>
            }
        </Panel>
    </main>;
};

export default UserAddressPage;