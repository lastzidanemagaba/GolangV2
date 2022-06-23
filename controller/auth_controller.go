package controller

import (
	"log"
	"net/http"
	"zidane/auth"
	"zidane/model"
	"zidane/responses"
	"zidane/service"

	"github.com/gin-gonic/gin"
	_ "golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var u model.UserLogin

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponses(http.StatusNotFound, 1, "Error", "Invalid JSON"))
		c.Abort()
		return
	}

	//check if the user exist:
	user, err := model.Model.GetUserByEmail(u.Email, u.Password)

	if err != nil {
		//c.JSON(http.StatusNotFound, err.Error())
		c.JSON(http.StatusNotFound, responses.ErrorResponses(http.StatusNotFound, 1, "Error", err.Error()))
		return
	}
	//since after the user logged out, we destroyed that record in the database so that same jwt token can't be used twice. We need to create the token again
	authData, err := model.Model.CreateAuth(user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponses(http.StatusNotFound, 1, "Error", err.Error()))
		return
	}
	var authD auth.AuthDetails
	authD.UserId = authData.UserID
	authD.AuthUuid = authData.AuthUUID
	token, loginErr := service.Authorize.SignIn(authD)
	if loginErr != nil {
		c.JSON(http.StatusForbidden, responses.ErrorResponses(http.StatusForbidden, 1, "Error", "Please try to login later"))
		return
	}
	type data_resp struct {
		ID                            uint64 `json:"id"`
		Device_Id                     string `json:"device_id"`
		Device_Name                   string `json:"device_name"`
		Email                         string `json:"email"`
		Phone                         string `json:"phone"`
		Password                      string `json:"password"`
		Temp_Password                 string `json:"temp_password"`
		Token                         string `json:"token"`
		Name                          string `json:"name"`
		Photo                         string `json:"photo"`
		Last_Juz                      string `json:"last_juz"`
		Last_Chapter                  string `json:"last_chapter"`
		Last_Verse                    string `json:"last_verse"`
		Last_Read_Date                string `json:"last_read_date"`
		Scheduled_Khatam_Date         string `json:"scheduled_khatam_date"`
		Default_Voice_Path            string `json:"default_voice_path"`
		Penyakit_Khusus               string `json:"penyakit_khusus"`
		Obat_Khusus                   string `json:"obat_khusus"`
		Gol_Darah                     string `json:"gol_darah"`
		Ktp_Image                     string `json:"ktp_image"`
		Passport_Image                string `json:"passport_image"`
		Vaccination_Certificate_Image string `json:"vaccination_certificate_image"`
		Alergi                        string `json:"alergi"`
		Verification_Token            string `json:"verification_token"`
		Verification_Token_Date       string `json:"verification_token_date"`
		Email_Verified                string `json:"email_verified"`
		Lat                           string `json:"lat"`
		Lng                           string `json:"lng"`
		Current_Location              string `json:"current_location"`
		Parking_Lat                   string `json:"parking_lat"`
		Parking_Lng                   string `json:"parking_lng"`
		Driver_Status                 string `json:"driver_status"`
		Google_Uid                    string `json:"google_uid"`
		Facebook_Uid                  string `json:"facebook_uid"`
		Fcm_Key                       string `json:"fcm_key"`
		Pushy_Token                   string `json:"pushy_token"`
		Panic_Marker_Id               string `json:"panic_marker_id"`
		Xmpp_Password                 string `json:"xmpp_password"`
		Lang_Code                     string `json:"lang_code"`
		Premium                       string `json:"premium"`
		Premium_Months                string `json:"premium_months"`
		Notification_Channel_Id       string `json:"notification_channel_id"`
		Group_Notification_Channel_Id string `json:"group_notification_channel_id"`
		Last_Premium_Date             string `json:"last_premium_date"`
	}
	res_data := data_resp{
		ID:                            user.ID,
		Email:                         user.Email,
		Device_Name:                   user.Device_Name,
		Password:                      u.Password,
		Temp_Password:                 user.Temp_Password,
		Name:                          user.Name,
		Photo:                         user.Photo,
		Last_Juz:                      user.Last_Juz,
		Last_Chapter:                  user.Last_Chapter,
		Last_Verse:                    user.Last_Verse,
		Last_Read_Date:                user.Last_Read_Date,
		Scheduled_Khatam_Date:         user.Scheduled_Khatam_Date,
		Default_Voice_Path:            user.Default_Voice_Path,
		Penyakit_Khusus:               user.Penyakit_Khusus,
		Obat_Khusus:                   user.Obat_Khusus,
		Gol_Darah:                     user.Gol_Darah,
		Ktp_Image:                     user.Ktp_Image,
		Passport_Image:                user.Passport_Image,
		Vaccination_Certificate_Image: user.Vaccination_Certificate_Image,
		Verification_Token:            user.Verification_Token,
		Verification_Token_Date:       user.Verification_Token_Date,
		Email_Verified:                user.Email_Verified,
		Lat:                           user.Lat,
		Lng:                           user.Lng,
		Current_Location:              user.Current_Location,
		Parking_Lat:                   user.Parking_Lat,
		Parking_Lng:                   user.Parking_Lng,
		Driver_Status:                 user.Driver_Status,
		Google_Uid:                    user.Google_Uid,
		Facebook_Uid:                  user.Facebook_Uid,
		Fcm_Key:                       user.Fcm_Key,
		Pushy_Token:                   user.Pushy_Token,
		Xmpp_Password:                 user.Xmpp_Password,
		Lang_Code:                     user.Lang_Code,
		Premium:                       user.Premium,
		Premium_Months:                user.Premium_Months,
		Notification_Channel_Id:       user.Notification_Channel_Id,
		Group_Notification_Channel_Id: user.Group_Notification_Channel_Id,
		Last_Premium_Date:             user.Last_Premium_Date,
		Token:                         token}

	c.JSON(http.StatusBadRequest, responses.SuccesResponses(http.StatusOK, 0, "Success", res_data))

}

func LogOut(c *gin.Context) {
	au, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", "Unauthorized"))
		return
	}
	delErr := model.Model.DeleteAuth(au)
	if delErr != nil {
		log.Println(delErr)
		c.JSON(http.StatusUnauthorized, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", "Unauthorized"))
		return
	}
	c.JSON(http.StatusBadRequest, responses.SuccesResponses(http.StatusOK, 0, "Success", "Successfully logged out"))
}
