
import keymirror from 'keymirror';

var orgActionTypes = keymirror({
  ORG_FETCH: null,
  ORGS_FETCH: null,
  ORG_RECEIVED: null,
  ORGS_RECEIVED: null
});

var spaceActionTypes = keymirror({
  SPACE_FETCH: null,
  SPACE_RECEIVED: null
});

var errorActionTypes = keymirror({
  FETCH: null
});

export { orgActionTypes, spaceActionTypes, errorActionTypes };