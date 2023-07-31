import {Store} from "./store";
import {createSelector} from "@reduxjs/toolkit";

const selectApp = ( store: Store ) => store.app;

export const selectProductCategories = createSelector( [selectApp], ( app ) => app.productCategories );

