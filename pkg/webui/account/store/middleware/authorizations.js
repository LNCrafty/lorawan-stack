// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import tts from '@account/api/tts'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as authorizations from '@account/store/actions/authorizations'

const getAuthorizationsLogic = createRequestLogic({
  type: authorizations.GET_AUTHORIZATIONS_LIST,
  process: async ({ action }) => {
    const {
      payload: { userId },
      meta: { selector },
    } = action
    const res = await tts.Authorizations.getAllAuthorizations(userId, selector)

    return { entities: res.authorizations, authorizationsTotalCount: res.totalCount }
  },
})

const deleteAuthorizationLogic = createRequestLogic({
  type: authorizations.DELETE_AUTHORIZATION,
  process: async ({ action }) => {
    const { userId, clientId } = action.payload

    return await tts.Authorizations.deleteAuthorization(userId, clientId)
  },
})

const getAuthorizationTokensLogic = createRequestLogic({
  type: authorizations.GET_AUTHORIZATION_TOKENS_LIST,
  process: async ({ action }) => {
    const {
      payload: { userId, clientId },
      meta: { selector },
    } = action
    const res = await tts.Authorizations.getAllTokens(userId, clientId, selector)

    return { entities: res.tokens, tokensTotalCount: res.totalCount }
  },
})

const deleteAuthorizationTokenLogic = createRequestLogic({
  type: authorizations.DELETE_AUTHORIZATION_TOKEN,
  process: async ({ action }) => {
    const { userId, clientId, id } = action.payload

    return await tts.Authorizations.deleteToken(userId, clientId, id)
  },
})

const deleteAllTokensLogic = createRequestLogic({
  type: authorizations.DELETE_ALL_TOKENS,
  process: async ({ action }) => {
    const { userId, clientId, ids } = action.payload

    return await Promise.all(ids.map(id => tts.Authorizations.deleteToken(userId, clientId, id)))
  },
})

export default [
  getAuthorizationsLogic,
  deleteAuthorizationLogic,
  getAuthorizationTokensLogic,
  deleteAuthorizationTokenLogic,
  deleteAllTokensLogic,
]
