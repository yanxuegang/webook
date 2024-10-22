// Copyright 2023 ecodeclub
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package domain

import "time"

const (
	BizProject = "project"
)

type Project struct {
	Id           int64
	SN           string
	Title        string
	Overview     string
	SystemDesign string
	// Github 仓库
	GithubRepo string
	// Gitee 仓库
	GiteeRepo string
	// 关联的八股文面试题集
	RefQuestionSet int64
	Status         ProjectStatus
	Desc           string
	Labels         []string
	Utime          int64
	// 其它字段
	Difficulties  []Difficulty
	Questions     []Question
	Resumes       []Resume
	Introductions []Introduction
	Combos        []Combo

	// 目前来说，我们只需要两个 SN，而不是需要维持整个 SPU
	// 后续如果需要 SPU 的其他字段，就重构为结构体
	// 作为产品售卖的 SPU
	ProductSPU string
	// 作为兑换码售卖的 SPU
	CodeSPU string
}

type ProjectStatus uint8

func (s ProjectStatus) ToUint8() uint8 {
	return uint8(s)
}

const (
	ProjectStatusUnknown ProjectStatus = iota
	ProjectStatusUnpublished
	ProjectStatusPublished
)

type Difficulty struct {
	Id       int64
	Title    string
	Analysis string
	Status   DifficultyStatus
	Utime    time.Time
	Content  string
}

type DifficultyStatus uint8

func (s DifficultyStatus) ToUint8() uint8 {
	return uint8(s)
}

const (
	DifficultyStatusUnknown DifficultyStatus = iota
	DifficultyStatusUnpublished
	DifficultyStatusPublished
)

type Resume struct {
	Id       int64
	Role     uint8
	Content  string
	Analysis string
	Status   ResumeStatus
	Utime    time.Time
}

type ResumeStatus uint8

func (s ResumeStatus) ToUint8() uint8 {
	return uint8(s)
}

const (
	ResumeStatusUnknown ResumeStatus = iota
	ResumeStatusUnpublished
	ResumeStatusPublished
)

type Introduction struct {
	Id       int64
	Role     uint8
	Content  string
	Analysis string
	Status   IntroductionStatus
	Utime    time.Time
}

type IntroductionStatus uint8

func (s IntroductionStatus) ToUint8() uint8 {
	return uint8(s)
}

const (
	IntroductionStatusUnknown IntroductionStatus = iota
	IntroductionStatusUnpublished
	IntroductionStatusPublished
)

// Question 虽然都叫做 Question
// 但是实际上，这个 Question 和 question 模块里面的是不同的
type Question struct {
	Id       int64
	Title    string
	Analysis string
	Answer   string
	Status   QuestionStatus
	Utime    time.Time
}

type QuestionStatus uint8

func (s QuestionStatus) ToUint8() uint8 {
	return uint8(s)
}

const (
	QuestionStatusUnknown QuestionStatus = iota
	QuestionStatusUnpublished
	QuestionStatusPublished
)

type Role uint8

const (
	RoleUnknown Role = iota
	RoleStudent
	RoleIntern
	RoleCoreDeveloper
	RoleManager
)

func (r Role) ToUint8() uint8 {
	return uint8(r)
}

// Combo 面试套路，连招
type Combo struct {
	Id      int64
	Title   string
	Content string
	Utime   int64
	Status  ComboStatus
}

type ComboStatus uint8

func (s ComboStatus) ToUint8() uint8 {
	return uint8(s)
}

const (
	ComboStatusUnknown ComboStatus = iota
	ComboStatusUnpublished
	ComboStatusPublished
)
