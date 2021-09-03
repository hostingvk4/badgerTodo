package models

func ToList(listDto ListDto) List {
	return List{Title: listDto.Title, Description: listDto.Description, UserId: listDto.UserId}
}

func ToListDto(list List) ListDto {
	return ListDto{ID: list.ID, Title: list.Title, Description: list.Description}
}

func ToListDTOs(lists []List) []ListDto {
	listDtos := make([]ListDto, len(lists))

	for i, val := range lists {
		listDtos[i] = ToListDto(val)
	}

	return listDtos
}
