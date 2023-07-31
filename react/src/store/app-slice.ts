import {Slice, SliceCaseReducers, createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {ProductCategory} from "../__generated__/graphql";
import apolloClientContext from "../common/contexts/apollo-client-context";
import {ALL_CATEGORIES_WITH_SPECIFICATIONS} from "../common/backend/graph/query/products";
import {AppState} from "./store";

export const getAllProductSpecifications = createAsyncThunk('APP', async (payload, thunkApi) => {
    return apolloClientContext.query({
        query: ALL_CATEGORIES_WITH_SPECIFICATIONS
    }).then( ( res ) => {
        if ( !res?.data?.productCategories ) {
            return thunkApi.rejectWithValue( [] );
        }

        return thunkApi.fulfillWithValue( res.data.productCategories );
    } ).catch( (err) => {
        return thunkApi.rejectWithValue( err );
    } );
});

const appSlice: Slice<AppState> = createSlice<AppState, SliceCaseReducers<AppState>, string>({
    name: 'app',
    initialState: {
        productCategories: []
    },
    reducers: {
    },
    extraReducers: (builder) => {
        builder
            .addCase(getAllProductSpecifications.rejected, (state, action) => {
                state.productCategories = [];
            })
            .addCase(getAllProductSpecifications.fulfilled, (state, action) => {
                state.productCategories = action.payload as ProductCategory[];
            });
    }
});

export default appSlice.reducer;