package services

import (
    "context"
    "errors"
    "github.com/Encedeus/panel/dto"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/role"
    "github.com/Encedeus/panel/ent/user"
    "github.com/google/uuid"
    "golang.org/x/exp/slices"
    "time"
)

func CreateUserRoleId(name string, email string, passwordHash string, roleId uuid.UUID) (*uuid.UUID, error) {

    roleData, err := Db.Role.Get(context.Background(), roleId)

    if err != nil {
        return nil, err
    }

    return CreateUser(name, email, passwordHash, roleData)
}

func CreateUserRoleName(name string, email string, passwordHash string, roleName string) (*uuid.UUID, error) {
    roleData, err := Db.Role.Query().Where(role.Name(roleName)).First(context.Background())

    if err != nil {
        return nil, err
    }

    return CreateUser(name, email, passwordHash, roleData)
}

func CreateUser(name string, email string, passwordHash string, role *ent.Role) (*uuid.UUID, error) {
    userData, err := Db.User.Create().
        SetName(name).
        SetEmail(email).
        SetPassword(passwordHash).
        SetRole(role).
        Save(context.Background())

    if err != nil {
        return nil, err
    }

    return &userData.ID, err
}

// DoesUserHavePermission checks if user's role have a permission
func DoesUserHavePermission(permission string, userID uuid.UUID) bool {
    userData, err := Db.User.Query().Where(user.IDEQ(userID)).Select("role_id").First(context.Background())
    if err != nil {
        return false
    }

    if IsUserDeleted(userData) {
        return false
    }

    roleData, err := Db.Role.Query().Where(role.ID(userData.RoleID)).Select("permissions").First(context.Background())
    if err != nil {
        return false
    }

    return slices.Contains(roleData.Permissions, permission) || slices.Contains(roleData.Permissions, "*")
}

// UpdateUser updates the user given an updateInfo dto
func UpdateUser(updateInfo dto.UpdateUserDTO) error {
    userData, err := Db.User.Query().Where(user.IDEQ(updateInfo.UserId)).First(context.Background())

    if err != nil {
        return err
    }

    if IsUserDeleted(userData) {
        return errors.New("user deleted")
    }

    if updateInfo.Name != "" {
        _, err = userData.Update().SetName(updateInfo.Name).Save(context.Background())
    }

    if updateInfo.Password != "" {
        _, err = userData.Update().SetPassword(updateInfo.Password).Save(context.Background())
    }

    if updateInfo.Email != "" {
        _, err = userData.Update().SetEmail(updateInfo.Email).Save(context.Background())
    }

    if updateInfo.RoleName != "" {
        roleData, roleErr := Db.Role.Query().Where(role.Name(updateInfo.RoleName)).First(context.Background())
        if roleErr != nil {
            return roleErr
        }
        _, err = userData.Update().SetRole(roleData).Save(context.Background())

    }

    if s := updateInfo.RoleId.String(); s != "" {
        roleData, roleErr := Db.Role.Query().Where(role.ID(updateInfo.RoleId)).First(context.Background())
        if roleErr != nil {
            return roleErr
        }
        _, err = userData.Update().SetRole(roleData).Save(context.Background())
    }

    return err
}

func DeleteUser(userId uuid.UUID) error {
    userData, err := Db.User.Query().Where(user.IDEQ(userId)).First(context.Background())
    if err != nil {
        return err
    }

    if IsUserDeleted(userData) {
        return errors.New("already deleted")
    }

    userData, err = userData.Update().SetDeletedAt(time.Now()).Save(context.Background())
    if err != nil {
        return err
    }

    return err
}

func GetUser(userId uuid.UUID) (*ent.User, error) {
    userData, err := Db.User.Query().
        Where(user.IDEQ(userId)).
        Select("uuid", "name", "created_at", "updated_at", "deleted_at", "email", "role_id").
        First(context.Background())
    if err != nil {
        return nil, err
    }

    if IsUserDeleted(userData) {
        return nil, errors.New("user deleted")
    }

    return userData, err
}

func DoesUserWithUUIDExist(userId uuid.UUID) bool {
    userData, err := Db.User.Query().Where(user.IDEQ(userId)).First(context.Background())

    if err != nil || IsUserDeleted(userData) {
        return false
    }

    return true
}

func GetLastUpdate(userId uuid.UUID) (int64, error) {
    userData, err := Db.User.Query().
        Where(user.IDEQ(userId), user.DeletedAtIsNil()).
        Select("updated_at").
        First(context.Background())

    if err != nil {
        return 0, nil
    }

    return userData.UpdatedAt.Unix(), nil
}
func IsUserDeleted(userData *ent.User) bool {
    return userData.DeletedAt.Unix() != -62135596800
}

func IsUserUpdated(userId uuid.UUID, issuedAt int64) (bool, error) {
    lastUpdate, err := GetLastUpdate(userId)
    if err != nil {
        return true, err
    }

    if lastUpdate > issuedAt {
        return true, nil
    }

    return false, nil
}