import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {
    MutationDeleteUserAddressArgs,
    UserAddress, UserAddressInput
} from "../../../__generated__/graphql";
import {ADD_USER_ADDRESS, DELETE_USER_ADDRESS, EDIT_USER_ADDRESS} from "../../../common/backend/graph/mutation/users";
import {USER_ADDRESSES} from "../../../common/backend/graph/query/users";
import apolloClientContext from "../../../common/contexts/apollo-client-context";
import UserState from "./state";

export type UserAddressRequest = {
    address_name: string
    address_line_1: string
    address_line_2: string
    "country[value]": string
    city: string
    province: string
    cap: string
}

export const fetchUserAddresses = createAsyncThunk('USER/FETCH_ADDRESSES', async (arg, thunkAPI) => {
    return apolloClientContext.query({
        query: USER_ADDRESSES
    }).then(res => {
        return thunkAPI.fulfillWithValue(res.data.userAddress);
    });
});

export const createUserAddress = createAsyncThunk('USER/CREATE_ADDRESS', async (arg: UserAddressRequest, thunkAPI) => {
    return apolloClientContext.mutate<any, { input: UserAddressInput }>({
        mutation: ADD_USER_ADDRESS,
        variables: {
            input: {
                addressName: arg.address_name,
                addressLine1: arg.address_line_1,
                addressLine2: arg.address_line_2,
                city: arg.city,
                country: arg["country[value]"],
                postalCode: arg.cap,
                province: arg.province,
                IsDefault: false
            }
        }
    }).then((res) => {
        if (!res.data) {
            return;
        }

        return thunkAPI.fulfillWithValue(res.data.createUserAddress);
    });
});

export const editUserAddress = createAsyncThunk( 'USER/EDIT_ADDRESS', async ( arg: { input: UserAddressRequest, id: string }, thunkAPI ) => {
    return apolloClientContext.mutate({
        mutation: EDIT_USER_ADDRESS,
        variables: {
            id: arg.id,
            input: {
                addressName: arg.input.address_name,
                addressLine1: arg.input.address_line_1,
                addressLine2: arg.input.address_line_2,
                city: arg.input.city,
                country: arg.input["country[value]"],
                postalCode: arg.input.cap,
                province: arg.input.province,
                IsDefault: false
            }
        }
    }).then( (res) => {
        return thunkAPI.fulfillWithValue( res.data.editUserAddress );
    } );
});

export const deleteUserAddress = createAsyncThunk('USER/DELETE_ADDRESS', async (arg: MutationDeleteUserAddressArgs, thunkAPI) => {
    return apolloClientContext.mutate({
        mutation: DELETE_USER_ADDRESS,
        variables: {
            id: arg.id
        }
    }).then(res => {
        return thunkAPI.fulfillWithValue(res.data.deleteUserAddress);
    });
});

const userSlice = createSlice<UserState, any, any>({
    name: "user",
    initialState: {
        addresses: []
    },
    reducers: {},
    extraReducers: builder => {
        builder.addCase(fetchUserAddresses.fulfilled, (state, action) => {
            state.addresses = action.payload as UserAddress[];
        });
        builder.addCase(createUserAddress.fulfilled, (state, action) => {
            state.addresses = action.payload as UserAddress[];
        });
        builder.addCase(deleteUserAddress.fulfilled, (state, action) => {
            state.addresses = action.payload as UserAddress[];
        });
        builder.addCase(editUserAddress.fulfilled, (state, action) => {
            state.addresses = action.payload as UserAddress[];
        });
    }
});


export default userSlice.reducer;