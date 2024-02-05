import face_recognition
import cv2
import numpy as np
import time

# Get a reference to webcam #0 (the default one)
video_capture = cv2.VideoCapture(0)

# Load a sample picture and learn how to recognize it.
pilatus_image = face_recognition.load_image_file("/Users/pilatus/workspaces/resources/pilatus.jpg")
pilatus_face_encoding = face_recognition.face_encodings(pilatus_image)[0]

# Load a second sample picture and learn how to recognize it.
biden_image = face_recognition.load_image_file("/Users/pilatus/workspaces/resources/biden.jpg")
biden_face_encoding = face_recognition.face_encodings(biden_image)[0]

# Create arrays of known face encodings and their names
known_face_encodings = [
    pilatus_face_encoding,
    biden_face_encoding
]
known_face_names = [
    "pilatus",
    "Joe Biden"
]


errCount = 0
succ = 0

while True:
    # Grab a single frame of video
    ret, frame = video_capture.read()
    rgb_frame = frame[:, :, ::-1]
    rgb_small_frame = np.ascontiguousarray(rgb_frame[:, :, ::-1])

    # Find all the faces and face enqcodings in the frame of video
    face_locations = face_recognition.face_locations(rgb_small_frame)
    try:
        face_encodings = face_recognition.face_encodings(rgb_small_frame, face_locations)
            # Loop through each face in this frame of video
        for (top, right, bottom, left), face_encoding in zip(face_locations, face_encodings):
            # See if the face is a match for the known face(s)
            matches = face_recognition.compare_faces(known_face_encodings, face_encoding)

            name = "Unknown"

            # If a match was found in known_face_encodings, just use the first one.
            if True in matches:
                first_match_index = matches.index(True)
                name = known_face_names[first_match_index]
                print(first_match_index, name)

            # Or instead, use the known face with the smallest distance to the new face
            face_distances = face_recognition.face_distance(known_face_encodings, face_encoding)
            best_match_index = np.argmin(face_distances)
            if matches[best_match_index]:
                name = known_face_names[best_match_index]
            # Draw a box around the face
            cv2.rectangle(frame, (left, top), (right, bottom), (0, 0, 255), 2)

            # Draw a label with a name below the face
            cv2.rectangle(frame, (left, bottom - 35), (right, bottom), (0, 0, 255), cv2.FILLED)
            font = cv2.FONT_HERSHEY_DUPLEX
            cv2.putText(frame, name, (left + 6, bottom - 6), font, 1.0, (255, 255, 255), 1)
            succ += 1
    except Exception as e:
        errCount += 1
        # print(e)
    # print(errCount, succ)
    cv2.imshow('Video', frame)

    cv2.waitKey(1)
    continue

time.sleep(10000)