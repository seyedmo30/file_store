package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"store/dto"
	"store/entity"
	"store/pkg/logs"
	"strings"
)

func (s Setup) RetrieveStore(ctx context.Context, query dto.RetrieveStoreRequest) (dto.RetrieveStoreResponse, error) {
	db := GetDB()
	queryName := ""
	if query.Name != nil {
		queryName = fmt.Sprintf("LOWER(name) = LOWER('%s') OR", *query.Name)
	}
	queryTag := fmt.Sprintf("ARRAY[ '%s' ]::varchar[] && tag or", strings.Join(query.Tag, "', '"))

	sql := fmt.Sprintf("SELECT id, name , type , hash , tag , filename FROM store_information WHERE %s %s false;", queryName, queryTag)
	logs.Connect().Error(sql)
	rows1, err := db.Query(sql)

	listStore := make([]entity.Store, 0, 5)
	if err != nil {
		logs.Connect().Error(err.Error())
	}
	defer rows1.Close()

	// Process the query results
	for rows1.Next() {
		var id, name, types, hash, tag, fileName string
		if err := rows1.Scan(&id, &name, &types, &hash, &tag, &fileName); err != nil {
			logs.Connect().Error(err.Error())
		}

		trimmedTag := strings.Trim(tag, "{}")

		tagSlice := strings.Split(trimmedTag, ",")

		for i, tag := range tagSlice {
			tagSlice[i] = strings.TrimSpace(tag)
		}

		logs.Connect().Debug(id + name + types + hash + tag)
		row := entity.Store{
			Name:     name,
			Hash:     hash,
			Tags:     tagSlice,
			Type:     types,
			FileName: fileName,
		}
		listStore = append(listStore, row)
	}

	if err := rows1.Err(); err != nil {
		logs.Connect().Error(err.Error())
	}

	return dto.RetrieveStoreResponse{Files: listStore}, nil
}

func (s Setup) CreateStore(ctx context.Context, request dto.CreateStoreRequest) error {

	db := GetDB()
	tagsArray := "{" + strings.Join(request.Tags, ",") + "}"
	sql := `INSERT INTO public.store_information (name, type, hash,filename, tag) VALUES ($1, $2, $3,$4 , $5) ;`
	_, err := db.Exec(sql,
		request.Name, request.Type, request.Hash, request.FileName, tagsArray)

	if err != nil {
		logs.Connect().Error(err.Error())
		return err
	}

	logs.Connect().Info("Data inserted successfully!")
	return nil
}

func (s Setup) DeleteStore(ctx context.Context, hashList []string) error {
	db := GetDB()
	if len(hashList) == 0 {
		return nil
	}

	quotedStrings := make([]string, len(hashList))

	for i, s := range hashList {
		quotedStrings[i] = fmt.Sprintf("'%s'", s)
	}

	// Join the quoted strings with commas
	joinedStrings := strings.Join(quotedStrings, ",")

	sql := fmt.Sprint("DELETE FROM store_information  WHERE hash IN (", joinedStrings, ") ;")
	logs.Connect().Info(sql)
	_, err := db.Exec(sql)

	if err != nil {
		logs.Connect().Error(err.Error())
		return err
	}
	logs.Connect().Info("Data deleted successfully!")
	return nil
}

func (s Setup) RetrieveFirstStore(ctx context.Context) (entity.Store, error) {
	db := GetDB()
	result := entity.Store{}
	rows := db.QueryRow("SELECT id, name , type , hash , tag , filename FROM store_information order by id limit 1;")

	var id, name, types, hash, tag, fileName string

	err := rows.Scan(&id, &name, &types, &hash, &tag, &fileName)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where no rows are returned
			logs.Connect().Info("No rows found.")
			return result, nil
		}
		logs.Connect().Error(err.Error())
		return result, err
	}

	trimmedTag := strings.Trim(tag, "{}")

	tagSlice := strings.Split(trimmedTag, ",")

	for i, tag := range tagSlice {
		tagSlice[i] = strings.TrimSpace(tag)
	}

	result = entity.Store{
		Name:     name,
		Hash:     hash,
		Tags:     tagSlice,
		Type:     types,
		FileName: fileName,
	}

	if err := rows.Err(); err != nil {
		logs.Connect().Error(err.Error())
	}

	return result, nil
}
