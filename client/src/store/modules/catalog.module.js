import CatalogService from '../../common/catalog.service'
import { FETCH_CATALOG, POST_CATALOG, DELETE_CATALOG, UPDATE_CATALOG } from '../actions.type'
import { SET_RESPONSE, SET_CATALOG, SET_DELETE_CATALOG, SET_UPDATE_CATALOG, ADD_CATALOG } from '../mutations.type'

const state = { 
    catalog: '',
    response: '',
}

const getters = {}

const actions = {
    async [FETCH_CATALOG]( {commit}) {
        const response =  CatalogService.fetchCatalog()
        const data = (await response).data
        commit(SET_CATALOG, data) 
    },
    async [POST_CATALOG]( { commit }, newCatalog ) {
       const response =  CatalogService.postCatalog(newCatalog)
        const data = (await response).data
        commit(ADD_CATALOG, data)
    },
    async [UPDATE_CATALOG] ({commit}, payload) {
        const response =  CatalogService.updateCatalog(payload)
        const data = (await response).data
        commit(SET_RESPONSE, data)
        commit(SET_UPDATE_CATALOG, payload)
    },
    async [DELETE_CATALOG] ({ commit }, id ) {
        const response =  CatalogService.deleteCatalog(id)
        const data = (await response).data
        commit(SET_RESPONSE, data)
        commit(SET_DELETE_CATALOG, id)
    }
}
const mutations = {
    [SET_CATALOG](state, catalog) {
        state.catalog = catalog
    },
    [SET_RESPONSE](state, response) {
        state.response = response
    },
    [ADD_CATALOG](state, catalog) {
        state.catalog.push(catalog)
    },
    [SET_UPDATE_CATALOG](state, newCatalog) {
        const t = state.catalog.find(x => x.id === newCatalog.id)
        t.name = newCatalog.name
    },
    [SET_DELETE_CATALOG] (state, id) {
        const catalog = state.catalog.find(x => x.id === id)
        const index = state.catalog.indexOf(catalog)
        state.catalog.splice(index,1)
    }
}

export default {
  state,
  getters,
  actions,
  mutations,
}
