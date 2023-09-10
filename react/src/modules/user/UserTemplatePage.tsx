import {QueryResult} from "@apollo/client";
import { faUser} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import {Img} from "react-image";
import {Outlet, useLoaderData, useNavigate} from "react-router-dom";
import {FetchUserQuery, UpdateAvatarMutation, UpdateAvatarMutationVariables} from "../../__generated__/graphql";
import {useMutation} from "../../common/backend/graph/hooks";
import {UPDATE_AVATAR} from "../../common/backend/graph/mutation/users";
import {useToast} from "../../store/toast";
import UserMenu from "./components/UserMenu";

const UserTemplatePage: React.FC = () => {
    const query = useLoaderData() as QueryResult<FetchUserQuery>;
    const navigate = useNavigate();
    const [mutate] = useMutation<UpdateAvatarMutation, UpdateAvatarMutationVariables>( UPDATE_AVATAR );
    const {error} = useToast();

    return <div className="h-full">
        <div className="flex flex-col md:flex-row items-stretch">
            <div className="p-4 flex flex-col bg-base-100 rounded-box m-1 shadow">
                    <div className="avatar circle placeholder aspect-square p-1 mx-auto">
                        <div className="bg-neutral-focus text-neutral-content rounded-full max-w-xs ring ring-base-300 ring-offset-base-100">
                            {
                                query.data?.user?.avatarUrl ?
                                    <Img className="bg-base-100 !object-contain" src={query.data?.user.avatarUrl}/> :
                                    <span className="text-6xl"> <FontAwesomeIcon className="w-36" icon={faUser}/> </span>
                            }
                        </div>
                    </div>
                <input
                    type="file"
                    accept=".png,.jpg,.jpeg,.svg,.webp"
                    className="file-input file-input-sm file-input-bordered w-full max-w-xs my-4"
                    onChange={ (e) => {
                    if ( !e.target.files ) {
                        return;
                    }

                    const file = e.target.files.item(0);

                    if ( !file ) {
                        return;
                    }

                    if (file && file.size > 1000000) {
                        error( "Il file selezionato e` troppo grande" );
                        return;
                    }

                    mutate( {
                        variables: {
                            userID: query?.data?.user?.id as string,
                            file: file
                        }
                    } ).then( () => {
                        navigate( 0 );
                    } );
                } }/>
                <UserMenu/>
            </div>
            <div className="w-full p-1">
                <Outlet/>
            </div>
        </div>
    </div>;
};

export default UserTemplatePage;