import {useQuery} from "@apollo/client";
import {faPencil, faPlus, faTimes} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import React from "react";
import {useTranslation} from "react-i18next";
import {useDispatch} from "react-redux";
import {useParams} from "react-router-dom";
import {UserAddress} from "../../__generated__/graphql";
import {USER_ADDRESSES} from "../../common/backend/graph/query/users";
import Button from "../../common/components-library/Button";
import Panel from "../../common/components-library/Panel";
import Table from "../../common/components-library/Table";
import Spinner from "../../common/components/Spinner";
import useModal from "../../hooks/useModal";
import ConfirmDeleteModal from "./components/ConfirmDeleteModal";
import UserAddressModal from "./components/UserAddressModal";
import {deleteUserAddress} from "./store/user-slice";
import {ThunkDispatch} from "redux-thunk";
import {useToast} from "../../store/toast";

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
    const {openModal, closeModal} = useModal();
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();

    if (loading) {
        return <Spinner/>;
    }

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
                                openModal(<UserAddressModal

                                    closeModal={() => {
                                        closeModal();
                                        refetch();
                                    }}
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
                                        .then(() => {
                                            toastr.success(t('user_address.delete_success'));
                                            refetch();
                                        })
                                        .catch(() => toastr.error(t('user_address.delete_error')));
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

    return <main className="h-full">
        <Panel>
            <div className="flex w-full justify-between">
                <h3 className="text-xl">{t('user_address.title_page')}</h3>
                <Button onClick={() => {
                    openModal(<UserAddressModal closeModal={() => {
                        closeModal();
                        refetch();
                    }}/>);
                }} type="primary"> <FontAwesomeIcon icon={faPlus}/> </Button>
            </div>
            <Table data={data.userAddress} columns={columns} hidePagination={true}/>
        </Panel>
    </main>;
};

export default UserAddressPage;