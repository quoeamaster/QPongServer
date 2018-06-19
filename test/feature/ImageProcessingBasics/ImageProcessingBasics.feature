Feature: Image Processing (Basic use cases)
    Image processing here includes use-case specific features only; hence
    the scenarios are focused on QPongServer's use-cases and not the use-cases
    for a generic Image processing library / framework

    Assumptions for the feature test:
    - image formats are jpg / jpeg / png ONLY

    Major use cases:
    - decode the format of the image file (determine is it jpg / png etc)
    - able to get the width and height of the image file
    - able to integrate with the "layers" of resources
        (could be plain text or other image resources) on top of the
        source image

    Scenario: 1) decode format of a jpg / jpeg image file PLUS the dimensions of the file
        Given there is an image at "/Users/jason.wong/ProjectData/Qpong/server/IMG_0062.JPG"
        Then after loading the image file; the format should be decoded into "jpg"
        And the width is "3024" and height is "4032" pixels

    Scenario: 2) decode format of a PNG image file PLUS the dimensions of the file
        Given there is an image at "/Users/jason.wong/ProjectData/Qpong/server/raindrops.png"
        Then after loading the image file; the format should be decoded into "png"
        And the width is "3000" and height is "1972" pixels