// MIT License
//
// Copyright (c) 2024 buuhuu
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package eventlog

var eventTemplate = map[EventId]WinEvent{
	System_App_Start: {
		Id:       System_App_Start,
		Message:  "Application started",
		Severity: Info,
		Type:     System,
	},
	System_App_Error: {
		Id:       System_App_Error,
		Message:  "Application encountered an error",
		Severity: Info,
		Type:     System,
	},
	System_App_Shutdown: {
		Id:       System_App_Shutdown,
		Message:  "Application shut down signal received",
		Severity: Warning,
		Type:     System,
	},
	System_App_Cleanup: {
		Id:       System_App_Cleanup,
		Message:  "Performing process cleanup",
		Severity: Warning,
		Type:     System,
	},
	System_Decoy_File_Created: {
		Id:       System_Decoy_File_Created,
		Message:  "Decoy file created",
		Severity: Info,
		Type:     System,
	},
	System_Decoy_File_Deleted: {
		Id:       System_Decoy_File_Deleted,
		Message:  "Decoy file deleted",
		Severity: Warning,
		Type:     System,
	},
}
