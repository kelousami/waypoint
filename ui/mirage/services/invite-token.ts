import { Token, NewTokenResponse } from 'waypoint-pb';

function createToken(): Token {
  let token = new Token();
  token.setTokenId(
    '3fwxJnSh32T9skH8NqseY8wuLQQynN6cnBYUCLTSxRJ6QCqLdEtUTY4hHjdDyHUiAarZC7WH1gZWypmQg8noi8ELfJxRe5131BFQWW3wzGW'
  );
  token.setInvite(true);
  return token;
}

export function create(schema: any, { params, requestHeaders }) {
  let resp = new NewTokenResponse();
  resp.setToken(createToken().getTokenId_asB64());
  return this.serialize(resp, 'application');
}
