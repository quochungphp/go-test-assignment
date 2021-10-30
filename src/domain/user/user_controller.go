package user

// type AuthController struct {
// 	AuthLoginAction AuthLoginAction
// }

// // name...
// func (ctrl AuthController) name() string {
// 	return "src.AuthController"
// }

// // Health ...
// func (ctrl AuthController) Login(w http.ResponseWriter, r *http.Request) {
// 	payload := &AuthLoginPayload{}
// 	body, err := ioutil.ReadAll(r.Body)
// 	defer r.Body.Close()

// 	if validErrs := payload.Validate(); len(validErrs) > 0 || err != nil {
// 		w.Header().Set("Content-type", "application/json")
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(validErrs)
// 		return
// 	}

// 	if err := json.Unmarshal(body, &payload); err != nil {
// 		panic("Error while unmarshalling")
// 	}

// 	_, err := ctrl.AuthLoginAction.Execute()
// 	if err != nil {
// 		w.Header().Set("Content-type", "application/json")
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode("Unauthorized")
// 		return
// 	}
// 	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	// w.Header().Set("X-Content-Type-Options", "nosniff")
// 	// w.WriteHeader(http.StatusOK)
// 	// json.NewEncoder(w).Encode("OK")
// }
